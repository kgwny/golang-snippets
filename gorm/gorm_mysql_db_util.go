package main

// gorm_mysql_db_util.go
// Utility to create and manage GORM + MySQL connections.

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewGormDB opens a GORM DB connection using the MySQL driver.
// dsn: MySQL DSN, e.g. "user:password@tcp(127.0.0.1:3306)/dbname?parseTime=true"
// maxOpenConns: maximum open connections
// maxIdleConns: maximum idle connections
// connMaxLifetime: maximum connection lifetime
// debug: when true sets GORM logger to Info, otherwise Silent
func NewGormDB(dsn string, maxOpenConns, maxIdleConns int, connMaxLifetime time.Duration, debug bool) (*gorm.DB, error) {
	logLevel := logger.Silent
	if debug {
		logLevel = logger.Info
	}
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm db: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	// connection pool settings
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	// quick ping to validate the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		// close underlying connections on error
		_ = sqlDB.Close()
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}

// CloseDB closes the underlying sql.DB from a gorm.DB
func CloseDB(d *gorm.DB) error {
	if d == nil {
		return nil
	}
	sqlDB, err := d.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}
	return sqlDB.Close()
}

// 呼び出し側のコーディング例
// example usage (in your main package):
/*
package main

import (
	"log"
	"time"

	"yourmodule/db"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/yourdb?parseTime=true"
	dbConn, err := db.NewGormDB(dsn, 20, 5, 30*time.Minute, true)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.CloseDB(dbConn)

	// use dbConn (type *gorm.DB)
}
*/
