package server

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"
	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"

	commentRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/repository"
	hotelRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/repository"

	"github.com/go-openapi/runtime/middleware"
	commentDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/delivery/http"
	commentUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/comment/usecase"
	hotelDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/delivery/http"
	hotelUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_main/internal/app/hotels/usecase"

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

func StartServer(db *sqlx.DB, log *logger.CustomLogger) {

	grpcSessionsConn, err := grpc.Dial(
		":8079",
		grpc.WithUnaryInterceptor(GetInterceptor(log)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcSessionsConn.Close()
	sessionService := sessionService.NewAuthorizationServiceClient(grpcSessionsConn)

	r := NewRouter()
	r.Methods("OPTIONS").Handler(middlewareApi.NewOptionsHandler())
	r.Use(middlewareApi.LoggerMiddleware(log))
	r.Use(middlewareApi.NewPanicMiddleware())
	r.Use(middlewareApi.MyCORSMethodMiddleware())

	repHot := hotelRepository.NewPostgresHotelRepository(db)
	repCom := commentRepository.NewCommentRepository(db)

	uHot := hotelUsecase.NewHotelUsecase(&repHot)
	uCom := commentUsecase.NewCommentUsecase(&repCom)

	sessMidleware := middlewareApi.NewSessionMiddleware(sessionService, u, log)
	csrfMidleware := middlewareApi.NewCsrfMiddleware(sessionService, log)
	r.Use(sessMidleware.SessionMiddleware())
	r.Use(csrfMidleware.CSRFCheck())

	hotelDelivery.NewHotelHandler(r, uHot, log)
	commentDelivery.NewCommentHandler(r, uCom, log)

	log.Info("Server started at port", configs.Port)
	//err := http.ListenAndServeTLS(configs.Port, "/etc/ssl/hostelscan.ru.crt", "/etc/ssl/hostelscan.ru.key", r)
	err = http.ListenAndServe(configs.Port, r)
	if err != nil {
		log.Error(err)
	}
}
