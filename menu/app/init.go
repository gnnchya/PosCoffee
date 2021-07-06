package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/menu/middleware"

	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/menu/app/user"
	userService "github.com/gnnchya/PosCoffee/menu/service/user"
)

type App struct {
	user *user.Controller
	middle middleware.Service
}

func New(userService userService.Service, middle middleware.Service) *App {
	return &App{
		user: user.New(userService),
		middle: middle,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	adminMiddleware := app.middle.Authorization(app.middle.Users)
	adminRoute := router.Group("/pos", adminMiddleware)
	{
		adminRoute.POST("/menu/:role", app.user.Create)
		adminRoute.PUT("/menu/:role", app.user.Update)
		adminRoute.DELETE("/menu/:id/:role", app.user.Delete)
		adminRoute.GET("/menu/:id/:role", app.user.Read)
		//adminRoute.GET("/menu", app.user.ReadAll)
		adminRoute.GET("/menu/search/:role", app.user.SearchMenu)
		adminRoute.GET("/menu/category/:role", app.user.SearchCategory)
		adminRoute.GET("/menu/ingredients/:role", app.user.SearchIngredient)
	}

	return app
}
