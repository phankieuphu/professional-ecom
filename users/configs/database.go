package config

import (
	"fmt"
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
			// when using docker, the database url is set in the env variable
			dsn := getEnv("DATABASE_URL", "")
			fmt.Println("Using DSN:", dsn)
			if dsn == "" {
				dbConfig := LoadDBConfig()
				dsn = dbConfig.DSN()
			}
			conn, err := gorm.Open(mysql.Open("app_user:app_pass@tcp(database:3306)/user_service?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
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
