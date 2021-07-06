package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/config"
	"log"
	"time"

	"github.com/gnnchya/PosCoffee/oAuth/app"
	consumerRepo "github.com/gnnchya/PosCoffee/oAuth/repository/mongodb"
	consumerService "github.com/gnnchya/PosCoffee/oAuth/service/consumer/implement"

	tokenRepo "github.com/gnnchya/PosCoffee/oAuth/repository/mongodb"
	tokenService "github.com/gnnchya/PosCoffee/oAuth/service/token/implement"

	oauthRepo "github.com/gnnchya/PosCoffee/oAuth/repository/oauth"
	grpcService "github.com/gnnchya/PosCoffee/oAuth/service/grpcClient/implement"

	validatorService "github.com/gnnchya/PosCoffee/oAuth/service/validator"
)

const (
	NETWORK = "tcp"
)
func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	cRepo, err := consumerRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBConsumerTableName)
	panicIfErr(err)
	tRepo, err := tokenRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTokenTableName)

	validator := validatorService.New(cRepo, tRepo)

	grpcRepo := oauthRepo.New()
	gService := grpcService.New(grpcRepo)
	panicIfErr(err)


	user := consumerService.New(validator, cRepo, gService)
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
