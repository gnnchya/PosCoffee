package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/oAuth/config"
	"github.com/gnnchya/PosCoffee/oAuth/service/util"
	"github.com/go-oauth2/oauth2/v4/manage"
	"log"
	"time"

	"github.com/gnnchya/PosCoffee/oAuth/app"
	consumerRepo "github.com/gnnchya/PosCoffee/oAuth/repository/mongodb"
	consumerService "github.com/gnnchya/PosCoffee/oAuth/service/consumer/implement"

	tokenRepo "github.com/gnnchya/PosCoffee/oAuth/repository/mongodb"
	tokenService "github.com/gnnchya/PosCoffee/oAuth/service/token/implement"

	oauthRepo "github.com/gnnchya/PosCoffee/oAuth/repository/oauth"
	grpcService "github.com/gnnchya/PosCoffee/oAuth/service/grpc"

	validatorService "github.com/gnnchya/PosCoffee/oAuth/service/validator"
)

const (
	NETWORK = "tcp"
)
func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	filter := util.NewFilters()

	cRepo, err := consumerRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBConsumerTableName)
	panicIfErr(err)
	tRepo, err := tokenRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBTokenTableName)
	panicIfErr(err)

	validator := validatorService.New(cRepo, tRepo)
	generateID, err := util.NewUUID()
	panicIfErr(err)
	uuid4 := util.NewUUID4()
	manager := manage.NewDefaultManager()

	oauthRepo := oauthRepo.New()

	consumer := consumerService.New(validator, cRepo, generateID, uuid4)

	token := tokenService.New(validator, manager, oauthRepo, tRepo, cRepo, filter, generateID)

	go grpcService.NewServer(appConfig, token, filter, tRepo)
	time.Sleep(1 * time.Second)

	return app.New(consumer, token)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

