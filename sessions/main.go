package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/internal/session/delivery"
	sessionsRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/internal/session/repository"
	sessionsUseCase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/internal/session/usecase"

	sessionService "github.com/go-park-mail-ru/2020_2_JMickhs/package/proto/sessions"

	csrfRepository "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/internal/csrf/repository"
	csrfUsecase "github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/internal/csrf/usecase"

	"google.golang.org/grpc"

	"github.com/go-park-mail-ru/2020_2_JMickhs/JMickhs_sessions/configs"
	"github.com/go-redis/redis/v8"
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
		log.Fatal(err)
	}
	fmt.Println(pong)
	return sessStore
}

func main() {
	configs.Init()
	store := NewSessStore()
	defer store.Close()

	repSes := sessionsRepository.NewSessionsUserRepository(store)
	repCsrf := csrfRepository.NewCsrfRepository(store)

	uSes := sessionsUseCase.NewSessionsUsecase(&repSes)
	uCsrf := csrfUsecase.NewCsrfUsecase(&repCsrf)

	server := grpc.NewServer()
	sessionService.RegisterAuthorizationServiceServer(server, delivery.NewSessionDelivery(uSes, uCsrf))

	listener, err := net.Listen("tcp", configs.SessionGrpcServicePort)
	if err != nil {
		log.Fatalf("can't listen port", err)
	}
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
