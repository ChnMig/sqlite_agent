package db

import (
	"database/sql"
	"time"

	"sqlite-agent/config"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var client *sql.DB

func GetClient() *sql.DB {
	if client == nil {
		Init()
	}
	return client
}

// Connect to the database
func Init() error {
	db, err := sql.Open("sqlite3", config.SqliteDSN)
	if err != nil {
		zap.L().Error("connect to sqlite failed", zap.Error(err))
		return err
	}

	// Set connection pool settings
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client = db
	return nil
}
