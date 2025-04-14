package config

import (
	"fmt"
	"os"
)

type EnvConfig struct {
	Database
}
type Database struct {
	DBHost     string
	DBName     string
	DBPort     string
	DBPassword string
	DBUser     string
}

func LoadDBConfig() *Database {
	return &Database{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "gateway_service"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBPassword: getEnv("DB_PASSWORD", "123456"),
		DBUser:     getEnv("DB_USER", "root"),
	}

}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}

func (db Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"your_user", db.DBPassword, db.DBHost, db.DBPort, db.DBName,
	)
}
