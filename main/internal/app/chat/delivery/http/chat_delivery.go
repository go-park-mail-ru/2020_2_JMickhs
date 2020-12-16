package chatDelivery

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/clientError"
	customerror "github.com/go-park-mail-ru/2020_2_JMickhs/package/error"

	tgbotapi "github.com/Syfaro/telegram-bot-api"

	uuid "github.com/satori/go.uuid"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	"github.com/spf13/viper"

	"github.com/gorilla/websocket"

	"github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHandler struct {
	ChatUseCase chat.Usecase
	bot         *tgbotapi.BotAPI
	roomTokens  map[string]string
	rooms       map[string]map[*connection]bool
	broadcast   chan message
	register    chan subscription
	unregister  chan subscription
	log         *logger.CustomLogger
}

type message struct {
	room    string
	message []byte
}
type connection struct {
	ws   *websocket.Conn
	send chan []byte
}

type subscription struct {
	conn *connection
	room string
}

func NewChatHandler(r *mux.Router, tgBot *tgbotapi.BotAPI, cu chat.Usecase, lg *logger.CustomLogger) *ChatHandler {
	handler := ChatHandler{
		ChatUseCase: cu,
		log:         lg,
		bot:         tgBot,
		roomTokens:  map[string]string{},
	}

	r.HandleFunc("/api/v1/ws/chat", handler.InitConnection).Methods("GET")
	r.Path("/api/v1/ws").Queries("chatID", "{chatID}", "token", "{token}").
		HandlerFunc(handler.InitConnectionForModer).Methods("GET")
	return &handler
}

// swagger:route GET /api/v1/ws/chat comment chat
// init chat
// responses:
//  200:
func (ch *ChatHandler) InitConnection(w http.ResponseWriter, r *http.Request) {
	//userID, ok := r.Context().Value(packageConfig.RequestUserID).(int)
	//if !ok {
	//	customerror.PostError(w, r, ch.log, errors.New("user unauthorized"), clientError.Unauthorizied)
	//	return
	//}

	chatID := uuid.NewV4().String()
	token := uuid.NewV4().String()
	query := r.URL.Query()

	query.Add("chatID", chatID)
	query.Add("token", token)
	str := "Ссылочка: " + r.URL.Scheme + r.Host + "/api/v1/ws" + "?" + query.Encode()

	config := tgbotapi.ChatConfig{ChatID: viper.GetInt64(configs.ConfigFields.ChatID)}
	chat, _ := ch.bot.GetChat(config)

	msg := tgbotapi.NewMessage(chat.ID, str)
	ch.bot.Send(msg)
	ch.roomTokens[chatID] = token
	ch.serveWs(w, r, chatID)
}

func (ch *ChatHandler) InitConnectionForModer(w http.ResponseWriter, r *http.Request) {
	chatID := mux.Vars(r)["chatID"]
	token := mux.Vars(r)["token"]
	if ch.roomTokens[chatID] != token {
		customerror.PostError(w, r, ch.log, errors.New("bad credentials"), clientError.Locked)
		return
	}
	ch.serveWs(w, r, chatID)
}

func (ch *ChatHandler) serveWs(w http.ResponseWriter, r *http.Request, roomID string) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ch.log.Error(err)
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, roomID}
	ch.register <- s
	go ch.read(&s)
	go ch.write(&s)
}

func (ch *ChatHandler) write(s *subscription) {
	conn := s.conn
	ticker := time.NewTicker(time.Duration(viper.GetInt64(configs.ConfigFields.PongWait)) * time.Minute)
	defer func() {
		conn.ws.Close()
		ticker.Stop()
	}()
	for {
		select {
		case msg, ok := <-conn.send:
			if !ok {
				conn.ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := conn.ws.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			if err := conn.ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (ch *ChatHandler) read(s *subscription) {
	conn := s.conn
	defer func() {
		conn.ws.Close()
	}()
	conn.ws.SetReadLimit(viper.GetInt64(configs.ConfigFields.MaxMessageSize))
	conn.ws.SetReadDeadline(time.Now().Add(time.Duration(viper.GetInt64(configs.ConfigFields.PongWait)) * time.Minute))
	conn.ws.SetPingHandler(func(string) error {
		conn.ws.SetReadDeadline(time.Now().Add(time.Duration(viper.GetInt64(configs.ConfigFields.PongWait)) * time.Minute))
		return nil
	})
	for {
		_, msg, err := conn.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				ch.log.Error(errors.New(""))
			}
		}
		m := message{room: s.room, message: msg}
		ch.broadcast <- m
	}

}

func (ch *ChatHandler) Run() {
	for {
		select {
		case s := <-ch.register:
			connections := ch.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				ch.rooms[s.room] = connections
			}
			ch.rooms[s.room][s.conn] = true
		case s := <-ch.unregister:
			connections := ch.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					if len(connections) == 0 {
						delete(ch.rooms, s.room)
					}
				}
			}
		case m := <-ch.broadcast:
			connections := ch.rooms[m.room]
			for c := range connections {
				select {
				case c.send <- m.message:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(ch.rooms, m.room)
					}
				}
			}

		}
	}
}
