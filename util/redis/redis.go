package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"golang-kafka/util/log"
	"os"
	"strconv"
)

var redisClient *redis.Client

func InitRedisClient(ctx context.Context) {
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Errorf("Error converting REDIS_DB to int: %v", err)
		return
	}

	if redisClient != nil {
		return
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Errorf("Redis connection was refused: %v", err)
		return
	}
}

func GetRedisClient() *redis.Client {
	if redisClient == nil {
		log.Errorf("Redis client is not initialized")
		return nil
	}

	return redisClient
}

func CloseRedisClient() {
	if redisClient == nil {
		log.Errorf("Redis client is nil, skipping close")
		return
	}

	err := redisClient.Close()
	if err != nil {
		log.Errorf("Redis connection close failed: %v", err)
	}
}
