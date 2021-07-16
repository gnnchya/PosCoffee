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
	verifyRoutes := router.Group("/verify")
	{
		verifyRoutes.GET("/verify_email/:token", app.user.VerifyEmail)
		verifyRoutes.POST("/forget_password", app.user.SendForgetPasswordUrl)
		verifyRoutes.POST("/change_password", app.user.ChangePassword)
		verifyRoutes.POST("/forget_password/:token", app.user.ForgetPassword)
	}

	loginMiddleware := app.middle.AuthorizationLogin(app.middle.Auth)
	loginRoute := router.Group("/user", loginMiddleware)
	{
		loginRoute.POST("/login", app.user.Login)
	}

	authMiddleware := app.middle.Authorization(app.middle.Auth)
	authRoute := router.Group("/auth", authMiddleware)
	{
		authRoute.GET("/list", app.user.List)
		authRoute.GET("/read/:id", app.user.Read)
		authRoute.POST("/logout", app.user.Logout)
		authRoute.PUT("/update", app.user.Update)
		authRoute.DELETE("/delete", app.user.Delete)

	}

	return app
}
