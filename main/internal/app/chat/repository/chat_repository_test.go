package chatRepository

import (
	"testing"
	"time"

	chat_model "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/models"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	redisServer *miniredis.Miniredis
	session     ChatRepository
	mock        sqlmock.Sqlmock
}

func (s *Suite) SetupSuite() {
	var err error
	s.redisServer, err = miniredis.Run()
	require.NoError(s.T(), err)

	addr := s.redisServer.Addr()
	redisConn := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	sqlxDb := sqlx.NewDb(db, "sqlmock")
	s.session = NewChatRepository(sqlxDb, redisConn)
	s.mock = mock
}

func (s *Suite) AfterTest(_, _ string) {
	s.SetupSuite()
}

func (s *Suite) TearDownSuite() {
	s.redisServer.Close()
}

func TestSessions(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestAddMessageInChat() {
	message := chat_model.Message{Message: "fsdfvcb", Room: "3", OwnerID: "1", Moderator: false}

	err := s.session.AddMessageInChat("3", message)
	require.NoError(s.T(), err)
	msg, _ := message.MarshalJSON()
	value, err := s.redisServer.RPush("3", string(msg))

	require.NoError(s.T(), err)
	require.Equal(s.T(), value, 1)
	s.redisServer.FastForward(time.Hour * 24)

	_, err = s.redisServer.List("3")
	require.Equal(s.T(), err, nil)

	s.redisServer.Close()
}

func (s *Suite) TestGetChatHistoryByID() {
	message := chat_model.Message{Message: "fsdfvcb", Room: "3", OwnerID: "1", Moderator: false}

	err := s.session.AddMessageInChat("3", message)
	require.NoError(s.T(), err)
	msg, _ := message.MarshalJSON()
	value, err := s.redisServer.RPush("3", string(msg))

	require.NoError(s.T(), err)
	require.Equal(s.T(), value, 1)
	s.redisServer.FastForward(time.Hour * 24)
	messages := []chat_model.Message{
		{Message: "fsdfvcb", Room: "3", OwnerID: "1", Moderator: false}}

	testMsgs, err := s.session.GetChatHistoryByID("3")

	require.NoError(s.T(), err)
	require.Equal(s.T(), testMsgs, messages)

	s.redisServer.Close()
}

func (s *Suite) TestAddOrGetChatByID() {

	message := chat_model.Message{Message: "fsdfvcb", Room: "3", OwnerID: "1", Moderator: false}

	err := s.session.AddMessageInChat("3", message)
	require.NoError(s.T(), err)
	msg, _ := message.MarshalJSON()
	value, err := s.redisServer.RPush("3", string(msg))

	require.NoError(s.T(), err)
	require.Equal(s.T(), value, 1)
	s.redisServer.FastForward(time.Hour * 24)
	messages := []chat_model.Message{
		{Message: "fsdfvcb", Room: "3", OwnerID: "1", Moderator: false}}

	rowChat := sqlmock.NewRows([]string{"chat_id"}).AddRow(
		3)
	s.mock.ExpectQuery(GetChatRequest).
		WithArgs(3).
		WillReturnRows(rowChat)

	testMsgs, err := s.session.AddOrGetChat("3", 3)

	require.NoError(s.T(), err)
	require.Equal(s.T(), testMsgs, messages)

	s.redisServer.Close()
}
