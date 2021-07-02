package app

import (
	// "touch/service/user"
	"github.com/gnnchya/PosCoffee/product/app/user"
	grpcService "github.com/gnnchya/PosCoffee/product/service/grpcClient"
	userService "github.com/gnnchya/PosCoffee/product/service/user"

	"github.com/gin-gonic/gin"
)

type App struct {
	user *user.Controller
	grpcService grpcService.Service
}

func New(userService userService.Service, grpcService grpcService.Service) *App {
	return &App{
		user: user.New(userService),
		grpcService: grpcService,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	apiRoutes := router.Group("/pos")
	{
		apiRoutes.POST("/product/transaction", app.user.Create)
		apiRoutes.GET("/product/transaction/:id", app.user.Read)
		apiRoutes.PUT("/product/transaction", app.user.Update)
		apiRoutes.DELETE("/product/transaction/:id", app.user.Delete)
		apiRoutes.GET("/product/transaction", app.user.ReadAll)
		apiRoutes.POST("/product/stock", app.user.CreateStock)
		apiRoutes.PUT("/product/stock", app.user.UpdateStock)
		apiRoutes.DELETE("/product/stock/:id", app.user.DeleteStock)
		apiRoutes.GET("/product/report", app.user.Report)
		apiRoutes.GET("/product/stock/:id", app.user.ReadStock)
		apiRoutes.GET("/product/stock/name/:name", app.user.ReadNameStock)
		apiRoutes.POST("/product/stock/category/:category", app.user.ReadCategoryStock)
		apiRoutes.POST("/product/reportSale", app.user.ReportSale)
		apiRoutes.POST("/product/reportStock", app.user.ReportStock)
	}

	return app
}
