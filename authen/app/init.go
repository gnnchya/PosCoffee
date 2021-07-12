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

type Roles struct{
	Owner string
	Staff string
	Admin string
	Report string
	Stock string
	Menu string
	Money string
	Transaction string
}

type App struct {
	user *user.Controller
	middle middleware.Service
	authService authentication.Service
}

var RoleID = Roles{"0001","0002","0003","0004","0005","0006","0007","0008"}

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
		apiRoutes.POST("/authen/staff", app.user.CreateStaff)
		apiRoutes.POST("/authen/owner", app.user.CreateOwner)
		apiRoutes.POST("/authen/admin", app.user.CreateAdmin)
		//apiRoutes.GET("/menu/:id", app.user.Read)
		//apiRoutes.PUT("/menu", app.user.Update)
		//apiRoutes.DELETE("/menu/:id", app.user.Delete)
		//apiRoutes.GET("/menu", app.user.ReadAll)
		//apiRoutes.POST("/menu/:id/cart", app.user.ToCart)
	}
	loginMiddleware := app.middle.AuthorizationLogin(app.middle.Auth)
	loginRoute := router.Group("/user", loginMiddleware)
	{
		loginRoute.POST("/login", app.user.Login)
	}

	//authMiddleware := app.middle.Authorization(app.middle.Auth)
	//authRoute := router.Group("/authorization", authMiddleware)
	//{
	//	authRoute.GET("/users", app.user.ReadAll)
	//}

	return app
}
