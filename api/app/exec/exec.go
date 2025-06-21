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

// Execute handles SQL operations.
func Execute(c *gin.Context) {
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
	response.ReturnData(c, rowsAffected)
}

// ExecuteQuery handles SQL SELECT operations.
func ExecuteQuery(c *gin.Context) {
	var req ExecuteSQLRequest
	if !middleware.CheckParam(&req, c) {
		return
	}

	db := db.GetClient()
	if db == nil {
		response.ReturnError(c, response.INTERNAL, "Database connection not initialized")
		return
	}

	rows, err := db.Query(req.SQL)
	if err != nil {
		response.ReturnError(c, response.INTERNAL, "Failed to execute SQL: "+err.Error())
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		response.ReturnError(c, response.INTERNAL, "Failed to retrieve columns: "+err.Error())
		return
	}

	results := []map[string]interface{}{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		if err := rows.Scan(scanArgs...); err != nil {
			response.ReturnError(c, response.INTERNAL, "Failed to scan row: "+err.Error())
			return
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = values[i]
		}
		results = append(results, rowMap)
	}

	response.ReturnData(c, results)
}
