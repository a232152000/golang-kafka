package kafka

import (
	"os"
	"strings"
)

type configKafka struct {
	BROKERS []string
}

var KafkaConfig configKafka

func KafkaBaseConfig() {
	KafkaConfig = configKafka{
		BROKERS: strings.Split(os.Getenv("KAFKA_BROKER"), ","),
	}
}
