package main

import (
	"context"
	setting "golang-kafka/config"
	"golang-kafka/util/kafka"
	"golang-kafka/util/log"
	notifier "golang-kafka/util/notify"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const KafkaNumConsumers = 3 // 要啟動的consumer worker 數量
const KafkaTopic = "jeff_test"
const KafkaGroup = "jeff_test"

func main() {
	defer kafka.CloseProducer()

	setting.InitConfig()

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < KafkaNumConsumers; i++ {
		wg.Add(1)
		go startConsumer(ctx, &wg, i)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-shutdown:
		log.Errorf("%v", sig)
		notifier.GetNotify().Send("worker consumer shutdown", "worker consumer shutdown")

		cancel()

		wg.Wait()
	}
}

func startConsumer(ctx context.Context, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Infof("Worker %d shutting down...", workerID)
			return
		default:
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Errorf("Worker %d panicked: %v. Restarting...", workerID, r)
						notifier.GetNotify().Send("Consumer panic", "Restarting...")
					}
				}()

				kafka.AsyncConsumer(ctx, KafkaTopic, KafkaGroup, &businessLogicA{})
			}()

			// Consumer 崩潰後等待 10 秒再重啟
			time.Sleep(10 * time.Second)
		}
	}
}
