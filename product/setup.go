package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/product/app"
	"github.com/gnnchya/PosCoffee/product/config"
	repoGrpc "github.com/gnnchya/PosCoffee/product/repository/grpc"
	"github.com/gnnchya/PosCoffee/product/repository/kafka"
	userRepo "github.com/gnnchya/PosCoffee/product/repository/user"
	grpcService "github.com/gnnchya/PosCoffee/product/service/grpc/implement"
	msgBrokerService "github.com/gnnchya/PosCoffee/product/service/msgbroker/implement"
	"github.com/gnnchya/PosCoffee/product/service/msgbroker/msgbrokerin"
	userService "github.com/gnnchya/PosCoffee/product/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/product/service/validator"
	"log"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName)
	uRepoMoney, err := userRepo.New2(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName, "THB")
	panicIfErr(err)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	validator := validatorService.New(uRepo)

	grpcRepo := repoGrpc.New(configGrpc(appConfig))
	user := userService.New(validator, uRepo, uRepoMoney, kRepo)

	//gService := grpcService.New(grpcRepo, user)

	msgService := msgBrokerService.New(kRepo, user)

	go grpcService.New(grpcRepo, user)

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

const (
	NETWORK = "tcp"
)

func configGrpc(appConfig *config.Config) *repoGrpc.Config {
	return &repoGrpc.Config{
		Network: NETWORK,
		Port:    appConfig.GRPCSenderHost,
	}
}
