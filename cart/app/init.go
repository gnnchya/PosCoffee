package app

import (
	// "touch/service/user"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/app/user"
	userService "github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/service/user"

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
	apiRoutes := router.Group("/api/v1")
	{
		apiRoutes.POST("/superheroes", app.user.Create)
		apiRoutes.PUT("/superheroes", app.user.Update)
		apiRoutes.GET("/superheroes/search", app.user.Search)
		apiRoutes.DELETE("/superheroes/:id", app.user.Delete)
		apiRoutes.GET("/superheroes/:id", app.user.View)
		apiRoutes.GET("/superheroes", app.user.ViewAll)
	}

	return app
}
