package main

import (
	"fmt"
	"os"

	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/public"
	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/settings"
	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/slack"
)

func main() {
	var err error
	var publicIPAddress string

	config := settings.FromEnv()
	if config.SlackWebHookURL == "" {
		fmt.Printf("SLACK_WEBHOOK_URL environment variable is not set\n")
		os.Exit(1)
	}

	fmt.Printf("The current machine's hostname is '%s'\n", config.Hostname)

	if publicIPAddress, err = public.IPAddress(config.IFConfigURL); err != nil {
		fmt.Printf("client: could not get public IP address: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Public IP address is: '%s'\n", publicIPAddress)
	message := `The public IP address of ` + config.Hostname + ` is: ` + publicIPAddress

	if err = slack.SendMessage(config.SlackWebHookURL, config.Hostname, message); err != nil {
		fmt.Printf("Failed to send slack a message: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("The message '%s' has been send to Slack OK\n", message)

	os.Exit(0)
}
