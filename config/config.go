package config

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"golang-kafka/util/database"
	"golang-kafka/util/kafka"
	"golang-kafka/util/log"
	notifier "golang-kafka/util/notify"
	"golang-kafka/util/redis"
)

type configKafka struct {
	BROKERS []string
}

var KafkaConfig configKafka

func InitConfig(ctx context.Context) {

	//env
	initEnv()
	//log
	log.InitLogger()
	//notify
	notifier.InitNotify()
	//Kafka Config
	kafka.KafkaBaseConfig()
	//init producer
	kafka.InitAsyncProducer()
	//init redis
	redis.InitRedisClient(ctx)

}

func initEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf(`load env error: %v`, err)
	}
}
