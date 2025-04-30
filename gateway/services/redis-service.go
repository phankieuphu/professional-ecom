package services

import (
	"context"
	"gateway/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	rdb *redis.Client
	ctx context.Context
}

var redisClient *redis.Client

func GetRDBClient() *RedisClient {
	if redisClient == nil {
		redisClient = redis.NewClient(config.LoadRedisConfig())
	}
	return &RedisClient{
		rdb: redisClient,
		ctx: context.Background(),
	}
}

func (c RedisClient) Get(key string) (string, error) {
	return c.rdb.Get(c.ctx, key).Result()
}

func (c RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return c.rdb.Set(c.ctx, key, value, expiration).Err()
}

func (c RedisClient) Del(key string) error {
	return c.rdb.Del(c.ctx, key).Err()
}

func (c RedisClient) Close() error {
	return c.rdb.Close()
}
