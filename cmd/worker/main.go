package main

import (
	"context"
	setting "golang-kafka/config"
	"golang-kafka/util/kafka"
	"golang-kafka/util/log"
	"golang-kafka/util/redis"
	"strconv"
	"sync"
)

const KakfaTopic = "jeff_test"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	setting.InitConfig(ctx)

	defer cancel()
	defer kafka.CloseProducer()
	defer redis.CloseRedisClient()

	wg := &sync.WaitGroup{} // 使用 WaitGroup 等待 goroutine 完成
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			//str := strconv.Itoa(int(time.Now().UnixNano()))
			str := strconv.Itoa(i)
			if err := kafka.ProduceMessage(KakfaTopic, str); err != nil {
				log.Errorf("Failed to send message from worker: %v\n", err)
			}
		}
	}()

	wg.Wait()
	//time.Sleep(5 * time.Second)
}
