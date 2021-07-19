package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/stock/config"
	"github.com/gnnchya/PosCoffee/stock/repository/grpc"
	"github.com/gnnchya/PosCoffee/stock/repository/kafka"
	msgBrokerService "github.com/gnnchya/PosCoffee/stock/service/msgbroker/implement"
	"github.com/gnnchya/PosCoffee/stock/service/msgbroker/msgbrokerin"
	"log"

	"github.com/gnnchya/PosCoffee/stock/app"
	userRepo "github.com/gnnchya/PosCoffee/stock/repository/user"
	grpcService "github.com/gnnchya/PosCoffee/stock/service/grpc/implement"
	userService "github.com/gnnchya/PosCoffee/stock/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/stock/service/validator"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName)
	panicIfErr(err)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	validator := validatorService.New(uRepo)

	grpcRepo := grpc.New(configGrpc(appConfig))
	grpcRepoReport := grpc.New(configGrpcRepo(appConfig))

	user := userService.New(validator, uRepo, kRepo)
	msgService := msgBrokerService.New(kRepo, user)
	//wg.Add(1)
	go grpcService.New(grpcRepo, user)
	go grpcService.New(grpcRepoReport, user)
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

func configGrpcRepo(appConfig *config.Config) *grpc.Config {
	return &grpc.Config{
		Network: "tcp",
		Port:    appConfig.GRPCSenderReportHost,
	}
}