package main

import (
	"context"
	"fmt"
	setting "golang-kafka/config"
	"golang-kafka/util/kafka"
	"golang-kafka/util/log"
	"golang-kafka/util/redis"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	setting.InitConfig(ctx)

	defer cancel()
	defer kafka.CloseProducer()
	defer redis.CloseRedisClient()

	redisTest := redis.GetRedisClient()
	if redisTest == nil {
		log.Errorf("Redis client is nil, exiting program")
	}

	err := redisTest.Set(ctx, "testKey", "testValue", 1*time.Hour).Err()
	if err != nil {
		fmt.Println("Failed to add testKey key-value pair")
		log.Errorf("%v", err)
	}

	val, _ := redisTest.Get(ctx, "testKey").Result()
	fmt.Println("redis:", val)
}
