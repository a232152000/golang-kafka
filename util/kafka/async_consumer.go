package kafka

import (
	"context"
	"golang-kafka/util/log"
	"strings"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

type BusinessLogic interface {
	DoBusiness(message *sarama.ConsumerMessage) error
}

type Consumer struct {
	ready chan bool
	logic BusinessLogic
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Errorf("message channel ConsumeClaim was closed")
				return nil
			}

			//先mark避免重複消化
			session.MarkMessage(message, "")

			err := consumer.logic.DoBusiness(message)
			if err != nil {
				log.Errorf("Error reason: %v, Message claimed: value =", err, string(message.Value))
				continue
			}

		case <-session.Context().Done():
			return nil
		}
	}
}

func AsyncConsumer(ctx context.Context, topic string, group string, logic BusinessLogic) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Group.Session.Timeout = 10 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 2 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup(KafkaConfig.BROKERS, group, config)
	if err != nil {
		log.Errorf("NewConsumerGroup err: %v", err)
		return
	}
	defer consumerGroup.Close()

	consumer := &Consumer{ready: make(chan bool), logic: logic}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			if err := consumerGroup.Consume(ctx, strings.Split(topic, ","), consumer); err != nil {
				log.Errorf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready

	wg.Wait()
}
