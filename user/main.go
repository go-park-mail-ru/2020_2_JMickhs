package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/grpcPackage"

	userGrpcDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/delivery/grpc"
	userHttpDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/delivery/http"

	userService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/user"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/middlewareUser"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2020_2_JMickhs/package/middlewareApi"
	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/configs"
	userRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/repository"
	userUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/usecase"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"google.golang.org/grpc"
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

func initRelativePath() string {
	_, fileName, _, _ := runtime.Caller(0)
	return filepath.ToSlash(filepath.Dir(filepath.Dir(fileName))) + "/"
}

func main() {
	validate := validator.New()

	configs.Init()
	db := InitDB()
	s3 := InitS3Session()

	configs.PrefixPath = initRelativePath()
	logOutput, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer logOutput.Close()
	log := logger.NewLogger(logOutput)

	grpcSessionsConn, err := grpc.Dial(
		configs.SessionGrpcServicePort,
		grpc.WithUnaryInterceptor(grpcPackage.GetInterceptor(log)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcSessionsConn.Close()
	sessionService := sessionService.NewAuthorizationServiceClient(grpcSessionsConn)

	r := mux.NewRouter()
	r.Methods("OPTIONS").Handler(middlewareApi.NewOptionsHandler())
	r.Use(middlewareApi.LoggerMiddleware(log))
	r.Use(middlewareApi.NewPanicMiddleware())
	r.Use(middlewareApi.MyCORSMethodMiddleware())

	rep := userRepository.NewPostgresUserRepository(db, s3)
	u := userUsecase.NewUserUsecase(&rep, validate)

	server := grpc.NewServer()
	userHttpDelivery.NewUserHandler(r, sessionService, u, log)

	sessMidleware := middlewareUser.NewSessionMiddleware(sessionService, u, log)
	csrfMidleware := middlewareApi.NewCsrfMiddleware(sessionService, log)

	r.Use(sessMidleware.SessionMiddleware())
	r.Use(csrfMidleware.CSRFCheck())

	userService.RegisterUserServiceServer(server, userGrpcDelivery.NewUserDelivery(u))

	listener, err := net.Listen("tcp", configs.UserGrpcServicePort)
	if err != nil {
		log.Fatalf("can't listen port", err)
	}
	go server.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server started at port", configs.UserHttpServicePort)
	//err = http.ListenAndServeTLS(configs.UserHttpServicePort, "/etc/ssl/hostelscan.ru.crt", "/etc/ssl/hostelscan.ru.key", r)
	err = http.ListenAndServe(configs.UserHttpServicePort, r)
	if err != nil {
		log.Error(err)
	}
}
