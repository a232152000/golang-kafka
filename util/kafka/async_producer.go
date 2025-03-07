package kafka

import (
	"golang-kafka/util/log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

var (
	producer sarama.AsyncProducer
	once     sync.Once
	err      error
	wg       sync.WaitGroup
)

func InitAsyncProducer() {
	once.Do(func() {
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 5
		config.Producer.Flush.Frequency = 2 * time.Second // 刷新頻率
		config.Producer.Flush.Messages = 10               // 緩存數量
		// 因為同步生產者在發送之後就必須返回狀態，所以需要兩個都返回
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true // 這個預設值就是 true 可以不用手動賦值

		producer, err = sarama.NewAsyncProducer(KafkaConfig.BROKERS, config)
		if err != nil {
			log.Errorf("NewAsyncProducer err: %v", err)
			return
		}

		go func() {
			for {
				select {
				case suc := <-producer.Successes():
					wg.Done()
					if suc != nil {
					}
				case err := <-producer.Errors():
					wg.Done()
					if err != nil {
						log.Errorf("Failed to Async produce message: %v, topic: %v, key: %v, value: %v", err.Err, err.Msg.Topic, err.Msg.Key, err.Msg.Value)
					}
				}
			}
		}()
	})
}

func getProducer() sarama.AsyncProducer {
	if producer == nil {
		log.Errorf("kafka async producer 尚未初始化")
	}
	return producer
}

func ProduceMessage(topic string, str string) error {
	producer := getProducer()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(str),
	}

	wg.Add(1)
	producer.Input() <- msg

	return nil
}

func CloseProducer() {
	wg.Wait()

	if producer != nil {
		producer.AsyncClose()
		producer = nil
	}
}
