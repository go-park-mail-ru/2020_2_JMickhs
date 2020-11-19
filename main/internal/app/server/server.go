package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"

	userDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/user/delivery/http"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"

	commentRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/repository"
	hotelRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/repository"
	userRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/user/repository"
	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/proto/sessions"

	"github.com/go-redis/redis/v8"

	"github.com/go-openapi/runtime/middleware"
	commentDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/delivery/http"
	commentUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/usecase"
	csrfRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/csrf/repository"
	csrfUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/csrf/usecase"
	hotelDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/delivery/http"
	hotelUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/usecase"
	middlewareApi "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/middleware"
	userUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/user/usecase"

	"google.golang.org/grpc"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/go-playground/validator/v10"

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
		Region:   aws.String(configs.S3Region),
		Endpoint: aws.String(configs.S3EndPoint),
	})))

}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./api/swagger")))

	return router
}

func GetInterceptor(log *logger.CustomLogger) func(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {

		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		log.Tracef("call=%v req=%#v reply=%#v time=%v err=%v",
			method, req, reply, time.Since(start), err)
		return err
	}
}

func StartServer(store *redis.Client, db *sqlx.DB, s3 *s3.S3, log *logger.CustomLogger) {
	validate := validator.New()

	grpcSessionsConn, err := grpc.Dial(
		":8079",
		grpc.WithUnaryInterceptor(GetInterceptor(log)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcSessionsConn.Close()
	sessionService := sessions.NewAuthorizationServiceClient(grpcSessionsConn)

	r := NewRouter()
	r.Methods("OPTIONS").Handler(middlewareApi.NewOptionsHandler())
	r.Use(middlewareApi.LoggerMiddleware(log))
	r.Use(middlewareApi.NewPanicMiddleware())
	r.Use(middlewareApi.MyCORSMethodMiddleware())

	rep := userRepository.NewPostgresUserRepository(db, s3)
	repHot := hotelRepository.NewPostgresHotelRepository(db)

	repCom := commentRepository.NewCommentRepository(db)
	repCsrf := csrfRepository.NewCsrfRepository(store)

	u := userUsecase.NewUserUsecase(&rep, validate)
	uHot := hotelUsecase.NewHotelUsecase(&repHot)
	uCom := commentUsecase.NewCommentUsecase(&repCom)
	uCsrf := csrfUsecase.NewCsrfUsecase(&repCsrf)

	sessMidleware := middlewareApi.NewSessionMiddleware(sessionService, u, log)
	csrfMidleware := middlewareApi.NewCsrfMiddleware(uCsrf, log)
	r.Use(sessMidleware.SessionMiddleware())
	r.Use(csrfMidleware.CSRFCheck())

	hotelDelivery.NewHotelHandler(r, uHot, log)
	userDelivery.NewUserHandler(r, sessionService, u, uCsrf, log)
	commentDelivery.NewCommentHandler(r, uCom, log)

	log.Info("Server started at port", configs.Port)
	//err := http.ListenAndServeTLS(configs.Port, "/etc/ssl/hostelscan.ru.crt", "/etc/ssl/hostelscan.ru.key", r)
	err = http.ListenAndServe(configs.Port, r)
	if err != nil {
		log.Error(err)
	}
}
