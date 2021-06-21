package app

import (
	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/product/app/user"
	userService "github.com/gnnchya/PosCoffee/product/service/user"

	"github.com/gin-gonic/gin"
)

type App struct {
	user *user.Controller
	// company *company.Controller
}

func New(userService userService.Service) *App {
	return &App{
		user: user.New(userService),
		// company: company.New(companyService),
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	apiRoutes := router.Group("/pos")
	{
		apiRoutes.POST("/product/transaction", app.user.Create)
		apiRoutes.GET("/product/transaction/:id", app.user.Read)
		apiRoutes.PUT("/product/transaction", app.user.Update)
		apiRoutes.DELETE("/product/transaction/:id", app.user.Delete)
	}

	return app
}
