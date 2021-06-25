package main

import (
	"github.com/gnnchya/PosCoffee/menu/app"
	"github.com/gnnchya/PosCoffee/menu/config"
	elasRepo "github.com/gnnchya/PosCoffee/menu/repository/elastic"
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


	validator := validatorService.New(elasRepo)
	user := userService.New(validator, elasRepo)
	return app.New(user)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
