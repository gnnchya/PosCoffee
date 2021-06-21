package main

import (
	"context"
	"github.com/gnnchya/PosCoffee/cart/config"
	elasRepo "github.com/gnnchya/PosCoffee/cart/repository/elastic"
	"github.com/gnnchya/PosCoffee/cart/repository/kafka"
	msgBrokerService "github.com/gnnchya/PosCoffee/cart/service/msgbroker/implement"
	"github.com/gnnchya/PosCoffee/cart/service/msgbroker/msgbrokerin"
	"log"

	"github.com/gnnchya/PosCoffee/cart/app"
	validatorService "github.com/gnnchya/PosCoffee/cart/service/validator"

	userRepo "github.com/gnnchya/PosCoffee/cart/repository/user"
	userService "github.com/gnnchya/PosCoffee/cart/service/user/implement"
)

func newApp(appConfig *config.Config) *app.App {
	ctx := context.Background()
	elasRepo, err := elasRepo.New(appConfig.ElasticDBEndpoint, appConfig.ElasticDBUsername, appConfig.ElasticDBPassword, "superhero")
	panicIfErr(err)
	uRepo, err := userRepo.New(ctx, appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBHeroTableName)
	panicIfErr(err)
	kRepo, err := kafka.New(configKafka(appConfig))
	panicIfErr(err)
	validator := validatorService.New(uRepo, elasRepo)

	user := userService.New(validator, uRepo, kRepo, elasRepo)
	msgService := msgBrokerService.New(kRepo, user)
	//wg.Add(1)
	msgService.Receiver(topics)
	//time.Sleep(10 * time.Second)
	return app.New(user)
}

func panicIfErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func configKafka(appConfig *config.Config) *kafka.Config {
	return &kafka.Config{
		BackOffTime:  appConfig.MessageBrokerBackOffTime,
		MaximumRetry: appConfig.MessageBrokerMaximumRetry,
		Host:         appConfig.MessageBrokerEndpoint,
		Group:        appConfig.MessageBrokerGroup,
		Version:      appConfig.MessageBrokerVersion,
	}
}
var topics = []msgbrokerin.TopicMsgBroker{
	msgbrokerin.TopicResponseCreate,
	msgbrokerin.TopicResponseUpdate,
	msgbrokerin.TopicResponseDelete,
}

