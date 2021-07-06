package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/app"
	"github.com/gnnchya/PosCoffee/menu/config"
	"github.com/gnnchya/PosCoffee/menu/middleware"
	elasRepo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
	redisRepo "github.com/gnnchya/PosCoffee/menu/repository/redis"
	userService "github.com/gnnchya/PosCoffee/menu/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/menu/service/validator"
	"log"
	"time"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()

	eRepo, err := elasRepo.New(appConfig.ElasticDBEndpoint, appConfig.ElasticDBUsername, appConfig.ElasticDBPassword, "menu")
	panicIfErr(err)
	rRepo, err := redisRepo.New(ctx, appConfig.RedisEndpoint, appConfig.RedisPassword, 24 * time.Hour)
	panicIfErr(err)

	validator := validatorService.New(eRepo)
	user := userService.New(validator, eRepo, rRepo)
	midService := middleware.New(user)
	return app.New(user, midService)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
