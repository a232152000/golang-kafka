package setting

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang-kafka/util/log"
	notifier "golang-kafka/util/notify"
)

func InitConfig() {

	//env
	initEnv()
	//log
	log.InitLogger()
	//notify
	notifier.InitNotify()
}

func initEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf(`load env error: %v`, err)
	}
}
