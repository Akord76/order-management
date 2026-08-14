package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDatabase opens a connection pool to the SQL Server instance
// described by cfg and stores it in the package-level DB variable.
func ConnectDatabase(cfg *AppConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("database: failed to connect: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("database: failed to get generic sql.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("database: ping failed: %v", err)
	}

	log.Println("database: connected to", cfg.DBName, "at", cfg.DBHost)
	DB = db
	return DB
}
