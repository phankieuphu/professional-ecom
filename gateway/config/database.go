package config

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func GetDb() *gorm.DB {
	once.Do(
		func() {
			dbConfig := LoadDBConfig()
			dsn := dbConfig.DSN()
			conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				log.Fatalf("failed to connect to DB: %v", err)
			}

			sqlDB, err := conn.DB()

			if err != nil {
				log.Fatalf("Failed to get generic DB from GORM")
			}
			// change to env default here
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(60 * time.Minute)
			sqlDB.SetConnMaxIdleTime(10 * time.Minute)

			dbInstance = conn

		})
	return dbInstance

}
