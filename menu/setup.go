package main

import (
	"github.com/gnnchya/PosCoffee/menu/config"
	elasRepo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
	"log"

	"github.com/gnnchya/PosCoffee/menu/app"
	validatorService "github.com/gnnchya/PosCoffee/menu/service/validator"

	userService "github.com/gnnchya/PosCoffee/menu/service/user/implement"
)

func newApp(appConfig *config.Config) *app.App {
	elasRepo, err := elasRepo.New(appConfig.ElasticDBEndpoint, appConfig.ElasticDBUsername, appConfig.ElasticDBPassword, "superhero")
	panicIfErr(err)

	validator := validatorService.New(elasRepo)

	user := userService.New(validator, elasRepo)

	return app.New(user)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}


