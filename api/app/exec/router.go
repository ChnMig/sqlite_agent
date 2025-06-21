package exec

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/sqlite/execute", ExecuteSQL)
}
