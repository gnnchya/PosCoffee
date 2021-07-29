package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gnnchya/PosCoffee/menu/config"
	//"github.com/gin-contrib/cors"
	cors "github.com/rs/cors/wrapper/gin"
	ginLogRus "github.com/toorop/gin-logrus"
	"net/http"
)

func main() {
	// Load config
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
	// Register route to gin
	_ = newApp(appConfig).RegisterRoute(router)

	// Gin start listen
	_ = router.Run()
}
