package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/config"
	"log"

	"github.com/gnnchya/PosCoffee/cart/app"
	validatorService "github.com/gnnchya/PosCoffee/cart/service/validator"

	userRepo "github.com/gnnchya/PosCoffee/cart/repository/user"
	userService "github.com/gnnchya/PosCoffee/cart/service/user/implement"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBHeroTableName)
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