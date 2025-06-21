package app

import (
	"github.com/gin-gonic/gin"

	"sqlite-agent/api/middleware"
)

func InitApi() *gin.Engine {
	// gin.Default uses Use by default. Two global middlewares are added, Logger(), Recovery(), Logger is to print logs, Recovery is panic and returns 500
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	router.SetTrustedProxies(nil)
	// Add consent cross-domain middleware
	router.Use(middleware.CorssDomainHandler())
	// api-v1
	// v1 := router.Group("/api/v1")
	// {
	// }
	return router
}
