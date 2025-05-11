package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/go-redis/redis/v8"
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

func LoadRedisConfig() *redis.Options {
	return &redis.Options{
		Addr:     getEnv("REDIS_HOST", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnvInt("REDIS_DB", 0),
	}
}

func LoadKafkaConsumerConfig() *kafka.ConfigMap {
	return &kafka.ConfigMap{
		"bootstrap.servers": GetKafkaBrokers(),
		"group.id":          GetKafkaGroup(),
		"auto.offset.reset": "earliest",
	}
}
func GetKafkaBrokers() string {
	return "localhost:9092"
}

func GetKafkaGroup() string {
	return "gateway_group"
}

func LoadListKafkaConsumerTopic() []string {
	return []string{"gateway"}
}

func LoadDBConfig() *Database {
	return &Database{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBName:     getEnv("DB_NAME", "user_service"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBUser:     getEnv("DB_USER", "root"),
	}

}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
func getEnvInt(key string, defaultVal int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func (db Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db.DBUser, db.DBPassword, db.DBHost, db.DBPort, db.DBName,
	)
}
