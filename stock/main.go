package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/stock/config"
	cors "github.com/rs/cors/wrapper/gin"
	ginLogRus "github.com/toorop/gin-logrus"
	"net/http"
)

func main() {
	appConfig := config.Get()
	log := setupLog()
	// Gin setup
	router := gin.New()
	router.Use(ginLogRus.Logger(log), gin.Recovery())
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000/"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete},
		ExposedHeaders:   []string{"X-Content-Length"},
		AllowedHeaders:   []string{""},
		AllowCredentials: false,
	}))

	// old
	router.Use(cors.AllowAll())
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run(":8083")
}
