package main

import (
	"github.com/gnnchya/PosCoffee/menu/app"
	"github.com/gnnchya/PosCoffee/menu/config"
	elasRepo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
	repoGrpc "github.com/gnnchya/PosCoffee/menu/repository/grpc"
	"github.com/gnnchya/PosCoffee/menu/service/grpc"
	grpcService "github.com/gnnchya/PosCoffee/menu/service/grpcClient/implement"
	userService "github.com/gnnchya/PosCoffee/menu/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/menu/service/validator"
	"log"
)
const (
	NETWORK = "tcp"
)
func newApp(appConfig *config.Config) *app.App {
	elasRepo, err := elasRepo.New(appConfig.ElasticDBEndpoint, appConfig.ElasticDBUsername, appConfig.ElasticDBPassword, "menu")
	panicIfErr(err)
	grpcRepo := repoGrpc.New(configGrpc(appConfig))
	gService := grpcService.New(grpcRepo)

	validator := validatorService.New(elasRepo)
	user := userService.New(validator, elasRepo, gService)
	go grpc.NewServer(appConfig,user)
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

