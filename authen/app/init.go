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
	user         *user.Controller
	middle       middleware.Service
	authService  authentication.Service
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
		apiRoutes.POST("/register/staff", app.user.CreateStaff)
		apiRoutes.POST("/register/owner", app.user.CreateOwner)
		apiRoutes.POST("/register/admin", app.user.CreateAdmin)
	}
	loginMiddleware := app.middle.AuthorizationLogin(app.middle.Auth)
	loginRoute := router.Group("/user", loginMiddleware)
	{
		loginRoute.POST("/login", app.user.Login)
	}

	//authMiddleware := app.middle.Authorization(app.middle.Auth)
	authRoute := router.Group("/authorization")
	{
		authRoute.POST("/verify/:token", app.user.VerifyEmail)
	}

	return app
}
