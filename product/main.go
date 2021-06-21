package main

import (
	"github.com/gin-gonic/gin"

	// "github.com/touchtechnologies-product/go-blueprint-clean-architecture/config"
	"github.com/gnnchya/InternAtTouch/tree/Develop-optimized/newApp/config"
)

func main() {
	// Load config
	appConfig := config.Get()

	// Init log format
	// log := setupLog()

	// Gin setup
	router := gin.New()

	// Set custom log for gin
	// router.Use(ginLogRus.Logger(log), gin.Recovery())

	// Jaeger setup
	// closer := setupJaeger(appConfig)
	// defer func() {
	// 	if err := closer.Close(); err != nil {
	// 		log.Error(err)
	// 	}
	// }()

	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run()
}
