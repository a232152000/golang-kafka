package main

import (
	setting "golang-kafka/config"
)

func main() {
	setting.InitConfig()

	//log用法
	//log.Errorf("message: %v ,message: %v. end", "messageA", "messageB")
	//notify用法
	//notifier.GetNotify().Send("title", "message")

}
