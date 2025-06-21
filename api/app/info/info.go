package info

import (
	"github.com/gin-gonic/gin"

	"sqlite-agent/api/response"
	"sqlite-agent/config"
	"sqlite-agent/db"
)

// SQLiteStatus 表示 SQLite 的运行状态
type SQLiteStatus struct {
	DatabasePath string `json:"database_path"`
	IsOpen       bool   `json:"is_open"`
	Error        string `json:"error,omitempty"`
}

// GetSQLiteStatus 获取 SQLite 的运行状态
func GetSQLiteStatus(c *gin.Context) {
	// 使用 db 包中的 GetClient 获取数据库连接
	db := db.GetClient()
	status := SQLiteStatus{
		DatabasePath: config.SqliteDSN,
		IsOpen:       db != nil,
	}

	if db == nil {
		status.Error = "Failed to connect to SQLite"
		response.ReturnErrorWithData(c, response.INTERNAL, status)
		return
	}

	// 测试数据库连接是否正常
	_, err := db.Exec("SELECT 1")
	if err != nil {
		status.Error = err.Error()
		response.ReturnErrorWithData(c, response.INTERNAL, status)
		return
	}

	response.ReturnData(c, status)
}
