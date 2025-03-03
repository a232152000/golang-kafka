package setting

import (
	"fmt"
	"github.com/joho/godotenv"
)

func InitConfig() {

	//env
	initEnv()
}

func initEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf(`load env error: %v`, err)
	}
}
