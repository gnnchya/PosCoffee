package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/menu/middleware"
	"github.com/gnnchya/PosCoffee/menu/service/grpcClient"

	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/menu/app/user"
	userService "github.com/gnnchya/PosCoffee/menu/service/user"
)

type App struct {
	user *user.Controller
	middle middleware.Service
	client grpcClient.Service
}

func New(userService userService.Service, middle middleware.Service, client grpcClient.Service) *App {
	return &App{
		user: user.New(userService),
		middle: middle,
		client: client,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	//adminMiddleware := app.middle.Authorization(app.middle.Users)
	//adminRoute := router.Group("/pos", adminMiddleware)
	adminRoute := router.Group("/pos")
	{
		adminRoute.POST("/menu", app.user.Create)
		adminRoute.PUT("/menu", app.user.Update)
		adminRoute.DELETE("/menu/:id", app.user.Delete)
		adminRoute.GET("/menu/:id", app.user.Read)
		adminRoute.GET("/menu", app.user.ReadAll)
		adminRoute.GET("/menu/search", app.user.SearchMenu)
		adminRoute.GET("/menu/category", app.user.SearchCategory)
		adminRoute.GET("/menu/ingredients", app.user.SearchIngredient)
	}

	return app
}
