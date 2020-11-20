package user

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"time"

	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/configs"
	"github.com/go-park-mail-ru/2020_2_JMickhs/package/logger"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"

	userDelivery "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/delivery/http"
	userRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/repository"
	userUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_user/internal/user/usecase"

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
		":8079",
		grpc.WithUnaryInterceptor(GetInterceptor(log)),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcSessionsConn.Close()
	sessionService := sessionService.NewAuthorizationServiceClient(grpcSessionsConn)

	rep := userRepository.NewPostgresUserRepository(db, s3)
	u := userUsecase.NewUserUsecase(&rep, validate)

	server := grpc.NewServer()
	userDelivery.NewUserHandler(r, sessionService, u, log)

	listener, err := net.Listen("tcp", ":8079")
	if err != nil {
		log.Fatalf("can't listen port", err)
	}
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
