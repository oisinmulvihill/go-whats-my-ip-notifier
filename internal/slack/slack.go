package slack

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendMessage(slackWebHookURL string, username string, message string) (err error) {
	client := &http.Client{}

	status := "good"

	body := []byte(`{
		"username": "` + username + `",
		"attachments": [
			{"color": "` + status + `", "text": "` + message + `"}
		]
	}`)

	r, err := http.NewRequest("POST", slackWebHookURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")

	res, err := client.Do(r)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Slack API called failed with status code: %d\n", res.StatusCode)
		return fmt.Errorf("failed to call slack webhook '%d'", res.StatusCode)
	}

	defer res.Body.Close()

	return nil
}
