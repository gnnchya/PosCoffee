package app

import (
	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/cart/app/user"
	userService "github.com/gnnchya/PosCoffee/cart/service/user"

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
		apiRoutes.POST("/cart", app.user.Create)
		apiRoutes.GET("/cart/:id", app.user.Read)
		apiRoutes.GET("/cart", app.user.ReadAll)
		apiRoutes.PUT("/cart", app.user.Update)
		apiRoutes.DELETE("/cart/:id", app.user.Delete)
		//apiRoutes.POST("/cart/:id/finish", app.user.Finish)
	}

	return app
}
