package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/config"
	"github.com/gnnchya/PosCoffee/product/repository/kafka"
	msgBrokerService "github.com/gnnchya/PosCoffee/product/service/msgbroker/implement"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	"log"

	"github.com/gnnchya/PosCoffee/product/app"
	validatorService "github.com/gnnchya/PosCoffee/product/service/validator"
	"github.com/gnnchya/PosCoffee/product/repository/grpc"
	userRepo "github.com/gnnchya/PosCoffee/product/repository/user"
	userService "github.com/gnnchya/PosCoffee/product/service/user/implement"
	grpcService"github.com/gnnchya/PosCoffee/product/service/grpc/implement"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName)
	uRepoMoney, err := userRepo.New2(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName, "THB")
	panicIfErr(err)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	validator := validatorService.New(uRepo)

	grpcRepo := grpc.New(configGrpc(appConfig))


	user := userService.New(validator, uRepo, uRepoMoney, kRepo)
	msgService := msgBrokerService.New(kRepo, user)
	go grpcService.New(grpcRepo, user)

	//wg.Add(1)
	msgService.Receiver(topics)
	//time.Sleep(10 * time.Second)
	return app.New(user)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func configKafka(appConfig *config.Config) *kafka.Config {
	return &kafka.Config{
		BackOffTime:  appConfig.MessageBrokerBackOffTime,
		MaximumRetry: appConfig.MessageBrokerMaximumRetry,
		Host:         appConfig.MessageBrokerEndpoint,
		Group:        appConfig.MessageBrokerGroup,
		Version:      appConfig.MessageBrokerVersion,
	}
}
var topics = []msgbrokerin.TopicMsgBroker{
	msgbrokerin.TopicResponseCreate,
	msgbrokerin.TopicResponseUpdate,
	msgbrokerin.TopicResponseDelete,
}

func configGrpc(appConfig *config.Config) *grpc.Config {
	return &grpc.Config{
		Network: "tcp",
		Port:    appConfig.GRPCSenderHost,
	}
}
