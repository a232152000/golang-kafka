package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang-kafka/util/kafka"
	"golang-kafka/util/log"
	notifier "golang-kafka/util/notify"
)

type configKafka struct {
	BROKERS []string
}

var KafkaConfig configKafka

func InitConfig() {

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
}

func initEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf(`load env error: %v`, err)
	}
}
