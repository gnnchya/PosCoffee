package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/app"
	"github.com/gnnchya/PosCoffee/menu/config"
	"github.com/gnnchya/PosCoffee/menu/middleware"
	elasRepo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
	repoGrpc "github.com/gnnchya/PosCoffee/menu/repository/grpc"
	redisRepo "github.com/gnnchya/PosCoffee/menu/repository/redis"
	grpcService "github.com/gnnchya/PosCoffee/menu/service/grpcClient/implement"
	userService "github.com/gnnchya/PosCoffee/menu/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/menu/service/validator"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	eRepo, err := elasRepo.New(appConfig.ElasticDBEndpoint, appConfig.ElasticDBUsername, appConfig.ElasticDBPassword, "menu")
	panicIfErr(err)
	rRepo, err := redisRepo.New(ctx, appConfig.RedisEndpoint, appConfig.RedisPassword, 24 * time.Hour)
	panicIfErr(err)
	grpcRepo := repoGrpc.New(configGrpcMiddleware(appConfig))
	gService := grpcService.New(grpcRepo)
	validator := validatorService.New(eRepo)
	user := userService.New(validator, eRepo, rRepo, gService)
	midService := middleware.New(user)
	return app.New(user, midService, gService)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
const (
	NETWORK = "tcp"
)
func configGrpcMiddleware(appConfig *config.Config) *repoGrpc.Config {
	return &repoGrpc.Config{
		Network: NETWORK,
		Port:    appConfig.GRPCAuthenHost,
	}
}
func setupLog() *logrus.Logger {
	lr := logrus.New()
	lr.SetFormatter(&logrus.JSONFormatter{})

	return lr
}