package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/config"
	"github.com/gnnchya/PosCoffee/stock/repository/kafka"
	msgBrokerService "github.com/gnnchya/PosCoffee/stock/service/msgbroker/implement"
	"github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"
	"log"

	"github.com/gnnchya/PosCoffee/stock/app"
	validatorService "github.com/gnnchya/PosCoffee/stock/service/validator"

	userRepo "github.com/gnnchya/PosCoffee/stock/repository/user"
	userService "github.com/gnnchya/PosCoffee/stock/service/user/implement"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName)
	panicIfErr(err)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	validator := validatorService.New(uRepo)

	grpcRepo := grpc.New()

	user := userService.New(validator, uRepo, kRepo)
	msgService := msgBrokerService.New(kRepo, user)
	//wg.Add(1)
	msgService.Receiver(topics)
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
	msgbrokerin.TopicCreate,
	msgbrokerin.TopicUpdate,
	msgbrokerin.TopicDelete,
}

func configGrpc(appConfig *config.Config) *grpc.Config {
	return &grpc.Config{
		Network: "tcp",
		Port:    appConfig.GRPCSenderHost,
	}
}