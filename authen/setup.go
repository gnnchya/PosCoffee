package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/app"
	"github.com/gnnchya/PosCoffee/authen/config"
	"github.com/gnnchya/PosCoffee/authen/middleware"
	userRepo "github.com/gnnchya/PosCoffee/authen/repository/mongodb"
	authenService "github.com/gnnchya/PosCoffee/authen/service/authentication/implement"
	"github.com/gnnchya/PosCoffee/authen/service/kafka"
	userService "github.com/gnnchya/PosCoffee/authen/service/user/implement"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	validatorService "github.com/gnnchya/PosCoffee/authen/service/validator"

	repoGrpc "github.com/gnnchya/PosCoffee/authen/repository/grpc"
	grpcService "github.com/gnnchya/PosCoffee/authen/service/grpcClient/implement"
	"log"
)
const (
	NETWORK = "tcp"
)
func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	filter := util.NewFilters()
	//TODO initiateDB in docker-compose
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName)
	grpcRepo := repoGrpc.New(configGrpc(appConfig))
	gService := grpcService.New(grpcRepo)
	panicIfErr(err)
	validator := validatorService.New(uRepo)
	kafkaRepo, err := kafka.New(configKafka(appConfig))
	user := userService.New(validator, uRepo, gService)
	auth := authenService.New(validator, appConfig, uRepo, filter, kafkaRepo)
	midService := middleware.New(auth, user)
	return app.New(user, midService, gService)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func configGrpc(appConfig *config.Config) *repoGrpc.Config {
	return &repoGrpc.Config{
		Network: NETWORK,
		Port:    appConfig.GRPCSenderHost,
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