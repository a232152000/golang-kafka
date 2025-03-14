package main

import (
	"context"
	"fmt"
	setting "golang-kafka/config"
	"golang-kafka/repository"
	"golang-kafka/util/database"
	"golang-kafka/util/kafka"
	"golang-kafka/util/redis"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	setting.InitConfig(ctx)

	defer cancel()
	defer kafka.CloseProducer()
	defer redis.CloseRedisClient()
	defer database.CloseDB()

	//不存在則建立table
	repository.CreateUser()

	//insert
	repository.InsertUser()

	//select
	users := repository.GetAllUser()
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", user.ID, user.Name, user.Email, user.Age)
	}
}
