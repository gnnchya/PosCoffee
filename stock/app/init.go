package app

import (
	"github.com/gnnchya/PosCoffee/stock/app/user"
	userService "github.com/gnnchya/PosCoffee/stock/service/user"

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
		apiRoutes.POST("/stock", app.user.Create)
		apiRoutes.PUT("/stock", app.user.Update)
		apiRoutes.GET("/stock/search", app.user.Search)
		apiRoutes.DELETE("/stock/:id", app.user.Delete)
		apiRoutes.GET("/stock/:id", app.user.Read)
		apiRoutes.GET("/stock/name/:name", app.user.ReadNameAll)
		apiRoutes.GET("/stock/category/:category", app.user.ReadCategoryAll)
	}

	return app
}
