package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/app"
	"github.com/gnnchya/PosCoffee/menu/config"
	elasRepo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
	repoGrpc "github.com/gnnchya/PosCoffee/menu/repository/grpc"
	redisRepo "github.com/gnnchya/PosCoffee/menu/repository/redis"
	grpcService "github.com/gnnchya/PosCoffee/menu/service/grpc/implement"
	userService "github.com/gnnchya/PosCoffee/menu/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/menu/service/validator"
	"log"
	"time"
)
const (
	NETWORK = "tcp"
)
func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	elasRepo, err := elasRepo.New(appConfig.ElasticDBEndpoint, appConfig.ElasticDBUsername, appConfig.ElasticDBPassword, "menu")
	panicIfErr(err)
	redisRepo, err := redisRepo.New(ctx, appConfig.RedisEndpoint, appConfig.RedisPassword, 24 * time.Hour)
	panicIfErr(err)
	grpcRepo := repoGrpc.New(configGrpc(appConfig))

	validator := validatorService.New(elasRepo)
	user := userService.New(validator, elasRepo, redisRepo)
	go grpcService.New(grpcRepo, user)
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
		Port:    appConfig.GRPCMenu,
	}
}