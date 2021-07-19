package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/oAuth/app/consumer"
	"github.com/gnnchya/PosCoffee/oAuth/app/token"
	consumerService "github.com/gnnchya/PosCoffee/oAuth/service/consumer"
	tokenService "github.com/gnnchya/PosCoffee/oAuth/service/token"
)

type App struct {
	consumer *consumer.Controller
	token *token.Controller
}

func New(consumerService consumerService.Service, tokenService tokenService.Service) *App {
	return &App{
		consumer: consumer.New(consumerService),
		token: token.New(tokenService),
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	apiRoutes := router.Group("/pos")
	{
		apiRoutes.POST("/consumer", app.consumer.Create)

		apiRoutes.POST("/request", app.token.Request)
		apiRoutes.DELETE("/revoke", app.token.RevokeToken)

		apiRoutes.POST("/validate", app.token.ValidateToken)
		apiRoutes.POST("/refresh", app.token.Refresh)
	}

	return app
}
