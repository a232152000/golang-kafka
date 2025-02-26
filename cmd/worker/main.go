package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("Go Version: %s\n", runtime.Version())

	for {
		time.Sleep(1 * time.Hour)
	}
}
