package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getPublicIPAddress(ifconfigURL string) (address string, err error) {

	client := &http.Client{}

	r, err := http.NewRequest("GET", ifconfigURL, nil)
	if err != nil {
		return string(""), err
	}

	res, err := client.Do(r)
	if err != nil {
		return string(""), err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return string(""), err
	}

	ipAddress := strings.TrimSpace(string(resBody))

	return ipAddress, nil
}

func sendNotification(slackWebHookURL string, username string, message string) (result string, err error) {
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
		return string(""), err
	}
	r.Header.Set("Content-Type", "application/json")

	res, err := client.Do(r)
	if err != nil {
		return string(""), err
	}
	if res.StatusCode != 200 {
		fmt.Printf("Slack API called failed with status code: %d\n", res.StatusCode)
		return string(""), fmt.Errorf("failed to call slack webhook '%d'", res.StatusCode)
	}

	defer res.Body.Close()

	return string("ok"), nil
}

func main() {
	ifconfigURL := os.Getenv("IFCONFIG_URL")
	if ifconfigURL == "" {
		ifconfigURL = "http://ifconfig.co"
	}

	slackWebHookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if slackWebHookURL == "" {
		fmt.Printf("SLACK_WEBHOOK_URL environment variable is not set\n")
		os.Exit(1)
	}

	hostname, err1 := os.Hostname()
	if err1 != nil {
		hostname = "<unknown hostname>"
	}
	fmt.Printf("The current machine's hostname is '%s'\n", hostname)

	publicIPAddress, err2 := getPublicIPAddress(ifconfigURL)
	if err2 != nil {
		fmt.Printf("client: could not get public IP address: %s\n", err2)
		os.Exit(1)
	}

	fmt.Printf("Public IP address is: '%s'\n", publicIPAddress)
	message := `The public IP address of ` + hostname + ` is: ` + publicIPAddress

	_, err3 := sendNotification(slackWebHookURL, hostname, message)
	if err3 != nil {
		fmt.Printf("Failed to send slack a message: %s\n", err3)
		os.Exit(1)
	}
	fmt.Printf("The message '%s' has been send to Slack OK\n", message)

	os.Exit(0)
}
