package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/go-playground/validator/v10"

	"github.com/go-park-mail-ru/2020_2_JMickhs/internal/pkg/logger"

	"github.com/go-park-mail-ru/2020_2_JMickhs/configs"
	commentDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/delivery/http"
	commentRepository "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/repository"
	commentUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/comment/usecase"

	sessionsRepository "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sessions/repository"
	sessionsUseCase "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/sessions/usecase"

	middlewareApi "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/middleware"

	"github.com/go-openapi/runtime/middleware"
	hotelDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/delivery/http"

	hotelUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/usecase"

	hotelRepository "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/hotels/repository"

	delivery "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/delivery/http"
	"github.com/go-redis/redis/v8"

	userRepository "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/repository"

	userUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/internal/app/user/usecase"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/mux"
)

func NewSessStore() *redis.Client {
	bd, _ := strconv.Atoi(configs.RedisConfig.Bd)
	sessStore := redis.NewClient(&redis.Options{
		Addr:     configs.RedisConfig.Address,
		Password: configs.RedisConfig.Password,
		DB:       bd, // use default DB
	})

	pong, err := sessStore.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pong)
	return sessStore
}

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
		Region:   aws.String("ru-msk"),
		Endpoint: aws.String("https://hb.bizmrg.com"),
	})))

}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	staticDir := "../static/"
	router.
		PathPrefix("/static").
		Handler(http.StripPrefix("/static", http.FileServer(http.Dir(staticDir))))

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./api/swagger")))

	return router
}


func StartServer(store *redis.Client,db *sqlx.DB,s3 *s3.S3,log *logger.CustomLogger) {

	validate := validator.New()

	r := NewRouter()
	r.Methods("OPTIONS").Handler(middlewareApi.NewOptionsHandler())
	r.Use(middlewareApi.MyCORSMethodMiddleware())
	r.Use(middlewareApi.NewPanicMiddleware())

	rep := userRepository.NewPostgresUserRepository(db)
	repHot := hotelRepository.NewPostgresHotelRepository(db)
	repSes := sessionsRepository.NewSessionsUserRepository(store)
	repCom := commentRepository.NewCommentRepository(db)

	u := userUsecase.NewUserUsecase(&rep, validate, s3)
	uHot := hotelUsecase.NewHotelUsecase(&repHot)
	uSes := sessionsUseCase.NewSessionsUsecase(&repSes)
	uCom := commentUsecase.NewCommentUsecase(&repCom)

	sessMidleware := middlewareApi.NewSessionMiddleware(uSes, u, log)
	r.Use(sessMidleware.SessionMiddleware())
	r.Use(middlewareApi.LoggerMiddleware(log))
	r.Use(middlewareApi.NewXssMiddleware())

	hotelDelivery.NewHotelHandler(r, uHot, log)
	delivery.NewUserHandler(r, uSes, u, log)
	commentDelivery.NewCommentHandler(r, uCom, log)

	log.Info("Server started at port", configs.Port)
	//err := http.ListenAndServeTLS(configs.Port, "/etc/ssl/hostelscan.ru.crt","/etc/ssl/hostelscan.ru.key",r)
	err := http.ListenAndServe(configs.Port,r)
	if err != nil {
		log.Error(err)
	}
}

