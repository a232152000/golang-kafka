package notifier

import (
	"bytes"
	"encoding/json"
	"golang-kafka/util/log"
	"net/http"
	"os"
)

type TeamsNotifier struct{}

type MessageCard struct {
	Type       string `json:"@type"`
	Context    string `json:"@context"`
	ThemeColor string `json:"themeColor"`
	Summary    string `json:"summary"`
	Sections   []struct {
		ActivityTitle string `json:"activityTitle"`
		Facts         []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"facts"`
		Markdown bool `json:"markdown"`
	} `json:"sections"`
}

func (t *TeamsNotifier) Send(title string, message string) error {
	msg := formatTeamsTemplate(title, message)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Errorf("Error notify teams json encode message: %v", err)
	}

	req, err := http.NewRequest("POST", os.Getenv("TEAMS_NOTIFY_URL"), bytes.NewBuffer(msgBytes))
	if err != nil {
		log.Errorf("Error notify teams json encode message: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error notify teams call teams api: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error notify teams status code not 200: %v", err)
	}

	return nil
}

func formatTeamsTemplate(title string, message string) MessageCard {
	messageCard := MessageCard{
		Type:       "MessageCard",
		Context:    "http://schema.org/extensions",
		ThemeColor: "d7000b",
		Summary:    title,
		Sections: []struct {
			ActivityTitle string `json:"activityTitle"`
			Facts         []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"facts"`
			Markdown bool `json:"markdown"`
		}{
			{
				ActivityTitle: title,
				Facts: []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				}{
					{Name: "log message：", Value: message},
				},
				Markdown: true,
			},
		},
	}

	return messageCard
}
