//  Golang service API for HotelScanner
//
//  Swagger spec.
//
//  Schemes: http
//  BasePath: /
//  Version: 1.0.0
//
//  Consumes:
//  - application/json
//  - multipart/form-data
//
//  Produces:
//	- application/json
//  swagger:meta
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

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
	_ "github.com/lib/pq"
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

func initDB() *sqlx.DB {
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

func initRelativePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	return filepath.ToSlash(filepath.Dir(filepath.Dir(fileName))) + "/"
}

func main() {
	configs.Init()
	store := NewSessStore()
	db := initDB()
	defer store.Close()

	configs.PrefixPath = initRelativePath()
	logOutput, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer logOutput.Close()

	log := logger.NewLogger(logOutput)

	r := NewRouter()
	r.Methods("OPTIONS").Handler(middlewareApi.NewOptionsHandler())
	r.Use(middlewareApi.MyCORSMethodMiddleware())
	r.Use(middlewareApi.NewPanicMiddleware())

	rep := userRepository.NewPostgresUserRepository(db)
	repHot := hotelRepository.NewPostgresHotelRepository(db)
	repSes := sessionsRepository.NewSessionsUserRepository(store)
	repCom := commentRepository.NewCommentRepository(db)

	u := userUsecase.NewUserUsecase(&rep)
	uHot := hotelUsecase.NewHotelUsecase(&repHot)
	uSes := sessionsUseCase.NewSessionsUsecase(&repSes)
	uCom := commentUsecase.NewCommentUsecase(&repCom)

	sessMidleware := middlewareApi.NewSessionMiddleware(uSes, u, log)
	r.Use(sessMidleware.SessionMiddleware())
	r.Use(middlewareApi.LoggerMiddleware(log))

	hotelDelivery.NewHotelHandler(r, uHot, log)
	delivery.NewUserHandler(r, uSes, u, log)
	commentDelivery.NewCommentHandler(r, uCom, log)
	log.Info("Server started at port", configs.Port)
	http.ListenAndServe(configs.Port, r)
}
