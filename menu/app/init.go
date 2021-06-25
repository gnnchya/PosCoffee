package app

import (
	"github.com/gin-gonic/gin"
	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/menu/app/user"
	grpcService "github.com/gnnchya/PosCoffee/menu/service/grpcClient"
	userService "github.com/gnnchya/PosCoffee/menu/service/user"
)

type App struct {
	user *user.Controller
	grpcService grpcService.Service
	// company *company.Controller
}

func New(userService userService.Service, grpcService grpcService.Service) *App {
	return &App{
		user: user.New(userService, grpcService),
		grpcService: grpcService,
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
	}

	return app
}
