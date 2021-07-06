package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/middleware"

	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/authen/app/user"
	grpcService "github.com/gnnchya/PosCoffee/authen/service/grpcClient"
	userService "github.com/gnnchya/PosCoffee/authen/service/user"
)

type App struct {
	user *user.Controller
	middle middleware.Service
	grpcService grpcService.Service

}

func New(userService userService.Service, middle middleware.Service, grpcService grpcService.Service) *App {
	return &App{
		user: user.New(userService),
		middle: middle,
		grpcService: grpcService,
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
