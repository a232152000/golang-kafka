package main

import (
	"github.com/IBM/sarama"
)

type businessLogicA struct{}

func (businessLogicA *businessLogicA) DoBusiness(message *sarama.ConsumerMessage) error {
	return nil
}
