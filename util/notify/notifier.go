package notifier

import (
	"fmt"
	"golang-kafka/util/log"
	"os"
)

var notifier Notifier

type Notifier interface {
	Send(title, message string) error
}

func NewNotifier() error {
	provider := os.Getenv("NOTIFY_PROVIDER")

	switch provider {
	case "teams":
		notifier = &TeamsNotifier{}
		return nil
	default:
		log.Errorf("notify unknown provider: %s", provider)
		return fmt.Errorf("unknown provider: %s", provider)
	}
}

func InitNotify() {
	err := NewNotifier()
	if err != nil {
		log.Errorf("Error creating notifier: %v\n", err)
		return
	}
}

func GetNotify() Notifier { return notifier }
