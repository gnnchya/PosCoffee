package main

import (
	"github.com/gin-gonic/gin"
	ginLogRus "github.com/toorop/gin-logrus"

	"github.com/gnnchya/PosCoffee/authorize/config"
	_ "github.com/gnnchya/PosCoffee/authorize/migrations"
)

func main() {
	// Load config
	appConfig := config.Get()

	// Init log format
	log := setupLog()

	// Migration database
	MigrateMongoDB(appConfig)

	// Gin setup
	router := gin.New()

	// Set custom log for gin
	router.Use(gin.Logger(), ginLogRus.Logger(log), gin.Recovery())

	// Jaeger setup
	closer := setupJaeger(appConfig)
	defer func() {
		if err := closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run()
}
