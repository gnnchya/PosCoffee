package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/authen/middleware"
	"github.com/gnnchya/PosCoffee/authen/service/authentication"
	//"github.com/gnnchya/PosCoffee/authorize/service/grpc/protobuf/authen"

	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/authen/app/user"
	userService "github.com/gnnchya/PosCoffee/authen/service/user"
)

type App struct {
	user *user.Controller
	middle middleware.Service
	authService authentication.Service
}

func New(userService userService.Service, authService authentication.Service, middle middleware.Service) *App {
	return &App{
		user: user.New(userService, authService),
		middle: middle,
		authService: authService,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	apiRoutes := router.Group("/pos")
	{
		apiRoutes.POST("/authen", app.user.Create)
		//apiRoutes.GET("/menu/:id", app.user.Read)
		//apiRoutes.PUT("/menu", app.user.Update)
		//apiRoutes.DELETE("/menu/:id", app.user.Delete)
		//apiRoutes.GET("/menu", app.user.ReadAll)
		//apiRoutes.POST("/menu/:id/cart", app.user.ToCart)
	}
	//loginMiddleware := app.middle.AuthorizationLogin(app.middle.Auth)
	//loginRoute := router.Group("/user", loginMiddleware)
	//{
	//	loginRoute.POST("/login", app.user.Create)
	//}
	//
	//authMiddleware := app.middle.Authorization(app.middle.Auth)
	//authRoute := router.Group("/authorization", authMiddleware)
	//{
	//	authRoute.GET("/users", app.user.ReadAll)
	//}

	return app
}
