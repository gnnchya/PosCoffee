package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/config"
	"log"
	"time"

	"github.com/gnnchya/PosCoffee/cart/app"
	repoGrpc "github.com/gnnchya/PosCoffee/cart/repository/grpc"
	userRepo "github.com/gnnchya/PosCoffee/cart/repository/user"
	grpcService "github.com/gnnchya/PosCoffee/cart/service/grpcClient/implement"
	userService "github.com/gnnchya/PosCoffee/cart/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/cart/service/validator"
)

const (
	NETWORK = "tcp"
)
func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTableName)
	grpcRepo := repoGrpc.New(configGrpc(appConfig))
	gService := grpcService.New(grpcRepo)
	panicIfErr(err)
	validator := validatorService.New(uRepo)

	user := userService.New(validator, uRepo, gService)
	time.Sleep(1 * time.Second)
	return app.New(user, gService)
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
