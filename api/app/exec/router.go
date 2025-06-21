package exec

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/insert", ExecuteInsert)
	router.POST("/delete", ExecuteDelete)
	router.POST("/update", ExecuteUpdate)
	router.POST("/query", ExecuteQuery)
}
