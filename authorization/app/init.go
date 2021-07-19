package app

import (
	"fmt"
	"net/http"

	"github.com/gnnchya/PosCoffee/authorize/config"
	"github.com/gnnchya/PosCoffee/authorize/docs"
	"github.com/gnnchya/PosCoffee/authorize/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gnnchya/PosCoffee/authorize/app/permissions"
	"github.com/gnnchya/PosCoffee/authorize/app/roles"
	permissionsService "github.com/gnnchya/PosCoffee/authorize/service/permissions"
	rolesService "github.com/gnnchya/PosCoffee/authorize/service/roles"
)

type App struct {
	permissions *permissions.Controller
	roles       *roles.Controller
	middle      middleware.Service
	swaggerHost string
	basePath    string
}

func New(permissionsService permissionsService.Service, rolesService rolesService.Service, conf *config.Config, middle middleware.Service) *App {
	return &App{
		permissions: permissions.New(permissionsService),
		roles:       roles.New(rolesService),
		middle:      middle,
		swaggerHost: conf.SwaggerHost,
		basePath:    conf.BasePath,
	}
}

func (app *App) RegisterRoute(router *gin.Engine) *App {
	docs.SwaggerInfo.Title = "Authorization Service API"
	docs.SwaggerInfo.Description = "API Spec Authorization Service."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = app.swaggerHost
	docs.SwaggerInfo.BasePath = app.basePath
	docPath := ginSwagger.URL(fmt.Sprintf("//%s/swagger/doc.json", app.swaggerHost))

	router.GET("/system/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
		return
	})

	//authMiddleware := app.middle.Authorization()
	authRoute := router.Group(docs.SwaggerInfo.BasePath)
	{
		authRoute.POST("/permissions", app.permissions.Create)
		authRoute.GET("/permissions/:id", app.permissions.Read)
		authRoute.PUT("/permissions/:id", app.permissions.Update)
		authRoute.DELETE("/permissions/:id", app.permissions.Delete)
		authRoute.GET("/permissions", app.permissions.List)

		authRoute.POST("/roles", app.roles.Create)
		authRoute.GET("/roles/:id", app.roles.Read)
		authRoute.PUT("/roles/:id", app.roles.Update)
		authRoute.DELETE("/roles/:id", app.roles.Delete)
		authRoute.GET("/roles", app.roles.List)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, docPath))

	return app
}
