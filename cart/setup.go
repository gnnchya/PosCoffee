package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/config"
	grpcService "github.com/gnnchya/PosCoffee/cart/service/grpcClient/implement"
	"log"

	"github.com/gnnchya/PosCoffee/cart/app"
	validatorService "github.com/gnnchya/PosCoffee/cart/service/validator"

	repoGrpc "github.com/gnnchya/PosCoffee/cart/repository/grpc"
	userRepo "github.com/gnnchya/PosCoffee/cart/repository/user"
	userService "github.com/gnnchya/PosCoffee/cart/service/user/implement"
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

	user := userService.New(validator, uRepo)
	//wg.Add(1)
	//time.Sleep(10 * time.Second)
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
