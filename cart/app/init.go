package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/cart/middleware"

	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/cart/app/user"
	grpcService "github.com/gnnchya/PosCoffee/cart/service/grpcClient"
	userService "github.com/gnnchya/PosCoffee/cart/service/user"
)

type App struct {
	user *user.Controller
	middle middleware.Service
	grpcService grpcService.Service
}

func New(userService userService.Service, grpcService grpcService.Service, middle middleware.Service) *App {
	return &App{
		user: user.New(userService, grpcService),
		middle: middle,
		grpcService: grpcService,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	//middleware := app.middle.Authorization(app.middle.Users)
	//apiRoutes := router.Group("/pos", middleware)
	apiRoutes := router.Group("/pos")
	{
		apiRoutes.POST("/cart", app.user.Create)
		apiRoutes.GET("/cart/:id", app.user.Read)
		apiRoutes.GET("/cart", app.user.ReadAll)
		apiRoutes.PUT("/cart", app.user.Update)
		apiRoutes.PUT("/cart/:id", app.user.AddToCart)
		apiRoutes.DELETE("/cart/:id", app.user.Delete)
		apiRoutes.GET("/cart/search", app.user.Search)
		apiRoutes.POST("/cart/:id/finish", app.user.Finish)
	}

	return app
}
