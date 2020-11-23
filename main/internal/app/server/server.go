package server

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/grpcPackage"

	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"
	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"log"
	"net/http"

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
		Region:   aws.String(viper.GetString(configs.ConfigFields.S3Region)),
		Endpoint: aws.String(viper.GetString(configs.ConfigFields.S3EndPoint)),
	})))

}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("../api/swagger")))

	return router
}

func StartServer(db *sqlx.DB, log *logger.CustomLogger) {

	grpcSessionsConn, err := grpc.Dial(
		viper.GetString(configs.ConfigFields.SessionGrpcServicePort),
		grpc.WithUnaryInterceptor(grpcPackage.GetInterceptor(log)),
		grpc.WithInsecure(),
	)
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
	r.Methods("OPTIONS").Handler(middlewareApi.NewOptionsHandler())
	r.Use(middlewareApi.LoggerMiddleware(log))
	r.Use(middlewareApi.NewPanicMiddleware())
	r.Use(middlewareApi.MyCORSMethodMiddleware())

	repHot := hotelRepository.NewPostgresHotelRepository(db)
	repCom := commentRepository.NewCommentRepository(db)

	uHot := hotelUsecase.NewHotelUsecase(&repHot, userService)
	uCom := commentUsecase.NewCommentUsecase(&repCom, userService)

	sessMidleware := middlewareApi.NewSessionMiddleware(sessionService, userService, log)
	csrfMidleware := middlewareApi.NewCsrfMiddleware(sessionService, log)
	r.Use(sessMidleware.SessionMiddleware())
	r.Use(csrfMidleware.CSRFCheck())

	hotelDelivery.NewHotelHandler(r, uHot, log)
	commentDelivery.NewCommentHandler(r, uCom, log)

	log.Info("Server started at port", viper.GetString(configs.ConfigFields.MainHttpServicePort))
	//err = http.ListenAndServeTLS(viper.GetString(configs.ConfigFields.MainHttpServicePort), "/etc/ssl/hostelscan.ru.crt", "/etc/ssl/hostelscan.ru.key", r)
	err = http.ListenAndServe(viper.GetString(configs.ConfigFields.MainHttpServicePort), r)
	if err != nil {
		log.Error(err)
	}
}
