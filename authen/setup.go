package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/app"
	"github.com/gnnchya/PosCoffee/authen/config"
	"github.com/gnnchya/PosCoffee/authen/middleware"
	repoGrpc "github.com/gnnchya/PosCoffee/authen/repository/grpc"
	userRepo "github.com/gnnchya/PosCoffee/authen/repository/mongodb"
	authenService "github.com/gnnchya/PosCoffee/authen/service/authentication/implement"
	grpcService "github.com/gnnchya/PosCoffee/authen/service/grpcClient/implement"
	userService "github.com/gnnchya/PosCoffee/authen/service/user/implement"
	"github.com/gnnchya/PosCoffee/authen/service/util"
	validatorService "github.com/gnnchya/PosCoffee/authen/service/validator"
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
	//kafkaRepo, err := kafka.New(configKafka(appConfig))
	user := userService.New(validator, uRepo, filter, gService)
	auth := authenService.New(validator, appConfig, uRepo, filter)
	midService := middleware.New(auth, user)
	return app.New(user, auth, midService)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func configGrpc(appConfig *config.Config) *repoGrpc.Config {
	return &repoGrpc.Config{
		Network: NETWORK,
		Port:    appConfig.GRPCAuthenHost,
	}
}
