package cache

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	if dbEnv := os.Getenv("REDIS_DB"); dbEnv != "" {
		if db, err := strconv.Atoi(dbEnv); err == nil {
			rdb.Options().DB = db
		}
	}

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to redis: %v", err))
	}

	return rdb
}
