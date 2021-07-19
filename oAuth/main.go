package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/oAuth/config"
)

func main() {
	// Load config
	appConfig := config.Get()

	// Gin setup
	router := gin.New()

	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run(":9090")
}
