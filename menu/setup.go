package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/menu/app"
	"github.com/gnnchya/PosCoffee/menu/config"
	elasRepo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
	redisRepo "github.com/gnnchya/PosCoffee/menu/repository/redis"
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

	validator := validatorService.New(elasRepo)
	user := userService.New(validator, elasRepo, redisRepo)
	return app.New(user)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
