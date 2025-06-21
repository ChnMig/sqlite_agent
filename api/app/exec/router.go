package exec

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("", Execute)
	router.POST("/query", ExecuteQuery)
}
