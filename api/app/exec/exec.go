package exec

import (
	"github.com/gin-gonic/gin"

	"sqlite-agent/api/middleware"
	"sqlite-agent/api/response"
	"sqlite-agent/db"
)

// ExecuteSQLRequest 表示执行 SQL 的请求结构
type ExecuteSQLRequest struct {
	SQL string `json:"sql" binding:"required"`
}

// ExecuteSQL 执行 SQL 语句
func ExecuteSQL(c *gin.Context) {
	var req ExecuteSQLRequest
	if !middleware.CheckParam(&req, c) {
		return
	}

	db := db.GetClient()
	if db == nil {
		response.ReturnError(c, response.INTERNAL, "Database connection not initialized")
		return
	}

	result, err := db.Exec(req.SQL)
	if err != nil {
		response.ReturnError(c, response.INTERNAL, "Failed to execute SQL: "+err.Error())
		return
	}

	rowsAffected, _ := result.RowsAffected()
	response.ReturnData(c, gin.H{"rows_affected": rowsAffected})
}
