package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/app"
	"github.com/gnnchya/PosCoffee/authen/config"
	userRepo "github.com/gnnchya/PosCoffee/authen/repository/mongodb"
	userService "github.com/gnnchya/PosCoffee/authen/service/user/implement"
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
	//TODO initiateDB in docker-compose
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName)
	grpcRepo := repoGrpc.New(configGrpc(appConfig))
	gService := grpcService.New(grpcRepo)
	panicIfErr(err)
	validator := validatorService.New(uRepo)
	user := userService.New(validator, uRepo, gService)
	return app.New(user)
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

