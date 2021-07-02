package app

import (
	"github.com/gin-gonic/gin"
	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/cart/app/user"
	grpcService "github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	userService "github.com/gnnchya/PosCoffee/cart/service/user"
)

type App struct {
	user *user.Controller
	grpcService grpcService.Service
}

func New(userService userService.Service, grpcService grpcService.Service) *App {
	return &App{
		user: user.New(userService, grpcService),
		grpcService: grpcService,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	apiRoutes := router.Group("/pos")
	{
		apiRoutes.POST("/cart", app.user.Create)
		apiRoutes.GET("/cart/:id", app.user.Read)
		//apiRoutes.GET("/cart", app.user.ReadAll)
		apiRoutes.PUT("/cart", app.user.Update)
		apiRoutes.DELETE("/cart/:id", app.user.Delete)
		apiRoutes.GET("/cart/search", app.user.Search)
		apiRoutes.GET("/cart/:id/finish", app.user.Finish)
	}

	return app
}
