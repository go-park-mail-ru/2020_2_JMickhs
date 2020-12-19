package server

import (
	"context"
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/Syfaro/telegram-bot-api"

	chatDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/delivery/http"

	chatRepository "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/repository"
	chatUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/chat/usecase"

	recommendRepository "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/repository"
	reccomendUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/recommendation/usecase"
	"github.com/go-redis/redis/v8"

	metrics2 "github.com/go-park-mail-ru/2020_2_JMickhs/package/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-park-mail-ru/2020_2_JMickhs/main/configs"
	commentDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/delivery/http"
	commentRepository "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/repository"
	commentUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/comment/usecase"
	hotelDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/delivery/http"
	hotelRepository "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/repository"
	hotelUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/hotels/usecase"
	wishlistDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/delivery/http"
	wishlistRepository "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/repository"
	wishlistUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/main/internal/app/wishlist/usecase"

	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/grpcPackage"

	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"
	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"log"
	"net/http"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"

	"google.golang.org/grpc"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
)

func InitDB() *sqlx.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configs.BdConfig.User,
		configs.BdConfig.Password,
		configs.BdConfig.DBName)

	fmt.Println(connStr)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func InitS3Session() *s3.S3 {
	return s3.New(session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(viper.GetString(configs.ConfigFields.S3Region)),
		Endpoint: aws.String(viper.GetString(configs.ConfigFields.S3EndPoint)),
	})))

}

func NewSessStore() *redis.Client {
	bd, _ := strconv.Atoi(configs.RedisConfig.Bd)
	sessStore := redis.NewClient(&redis.Options{
		Addr:     configs.RedisConfig.Address,
		Password: configs.RedisConfig.Password,
		DB:       bd, // use default DB
	})

	pong, err := sessStore.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)
	return sessStore
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./api/swagger")))

	return router
}

func StartServer(db *sqlx.DB, log *logger.CustomLogger, s3 *s3.S3) {

	grpcSessionsConn, err := grpc.Dial(
		viper.GetString(configs.ConfigFields.SessionGrpcServicePort),
		grpc.WithUnaryInterceptor(grpcPackage.GetInterceptor(log)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcSessionsConn.Close()

	sessionService := sessionService.NewAuthorizationServiceClient(grpcSessionsConn)

	grpcUserConn, err := grpc.Dial(
		viper.GetString(configs.ConfigFields.UserGrpcServicePort),
		grpc.WithUnaryInterceptor(grpcPackage.GetInterceptor(log)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}

	userService := userService.NewUserServiceClient(grpcUserConn)

	r := NewRouter()
	metrics := metrics2.RegisterMetrics()
	r.Methods("OPTIONS").Handler(middlewareApi.NewOptionsHandler())
	r.Handle("/api/v1/metrics", promhttp.Handler())
	r.Use(middlewareApi.LoggerMiddleware(log, metrics))
	r.Use(middlewareApi.NewPanicMiddleware(metrics))
	r.Use(middlewareApi.MyCORSMethodMiddleware())

	repHot := hotelRepository.NewPostgresHotelRepository(db, s3)
	repCom := commentRepository.NewCommentRepository(db, s3)
	repWish := wishlistRepository.NewPostgreWishlistRepository(db)
	repChat := chatRepository.NewChatRepository(db)
	store := NewSessStore()
	defer store.Close()

	repRecommendation := recommendRepository.NewPostgreRecommendationRepository(db, store)

	uChat := chatUsecase.NewChatUseCase(&repChat)
	uHot := hotelUsecase.NewHotelUsecase(&repHot, userService, &repWish)
	uCom := commentUsecase.NewCommentUsecase(&repCom, userService)
	uWish := wishlistUsecase.NewWishlistUseCase(&repWish, &repHot)
	uRecommendation := reccomendUsecase.NewRecommendationsUseCase(&repRecommendation)

	sessMidleware := middlewareApi.NewSessionMiddleware(sessionService, userService, log)
	csrfMidleware := middlewareApi.NewCsrfMiddleware(sessionService, log)
	r.Use(sessMidleware.SessionMiddleware())
	r.Use(csrfMidleware.CSRFCheck())

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BotToken"))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	hotelDelivery.NewHotelHandler(r, uHot, uRecommendation, log)
	commentDelivery.NewCommentHandler(r, uCom, log)
	wishlistDelivery.NewWishlistHandler(r, uWish, uHot, log)
	chatHandler := chatDelivery.NewChatHandler(r, bot, uChat, log)
	go chatHandler.Run()
	err = http.ListenAndServeTLS(viper.GetString(configs.ConfigFields.MainHttpServicePort), "/etc/ssl/hostelscan/hostelscan.ru.crt", "/etc/ssl/hostelscan/hostelscan.ru.key", r)
	//err = http.ListenAndServe(viper.GetString(configs.ConfigFields.MainHttpServicePort), r)
	if err != nil {
		log.Error(err)
	}
	log.Info("Server started at port", viper.GetString(configs.ConfigFields.MainHttpServicePort))
}
