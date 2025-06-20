package db

import (
	"time"

	"sqlite-agent/config"

	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var client *gorm.DB

func GetClient() *gorm.DB {
	if client == nil {
		Init()
	}
	return client
}

// Connect to the database
func Init() error {
	db, err := gorm.Open(sqlite.Open(config.SqliteDSN), &gorm.Config{})
	if err != nil {
		zap.L().Error("connect to sqlite failed", zap.Error(err))
		return err
	}
	pgDB, err := db.DB()
	if err != nil {
		zap.L().Error("get db failed", zap.Error(err))
		return err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	pgDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	pgDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	pgDB.SetConnMaxLifetime(time.Hour)
	client = db
	return nil
}
