package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/authen/app"
	"github.com/gnnchya/PosCoffee/authen/config"
	elasRepo "github.com/gnnchya/PosCoffee/authen/repository/elastic"
	redisRepo "github.com/gnnchya/PosCoffee/authen/repository/redis"
	userService "github.com/gnnchya/PosCoffee/authen/service/user/implement"
	validatorService "github.com/gnnchya/PosCoffee/authen/service/validator"
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
