package main

import (
	setting "golang-kafka/config"
	"golang-kafka/util/kafka"
	"golang-kafka/util/log"
	"strconv"
	"sync"
)

const KakfaTopic = "jeff_test"

func main() {
	defer kafka.CloseProducer()

	setting.InitConfig()
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
