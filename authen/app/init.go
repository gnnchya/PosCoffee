package app

import (
	"github.com/gin-gonic/gin"
	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/authen/app/user"
	userService "github.com/gnnchya/PosCoffee/authen/service/user"
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
		apiRoutes.POST("/menu", app.user.Create)
		apiRoutes.GET("/menu/:id", app.user.Read)
		apiRoutes.PUT("/menu", app.user.Update)
		apiRoutes.DELETE("/menu/:id", app.user.Delete)
		apiRoutes.GET("/menu", app.user.ReadAll)
		//apiRoutes.POST("/menu/:id/cart", app.user.ToCart)
	}

	return app
}
