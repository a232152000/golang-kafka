package main

import (
	setting "golang-kafka/config"
	"time"
)

func main() {
	setting.InitConfig()

	for {
		time.Sleep(1 * time.Hour)
	}
}
