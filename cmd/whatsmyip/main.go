package main

import (
	"fmt"
	"os"

	"internal/public"
	"internal/slack"
)

func main() {
	var err error
	var hostname string
	var publicIPAddress string

	ifconfigURL := os.Getenv("IFCONFIG_URL")
	if ifconfigURL == "" {
		ifconfigURL = "http://ifconfig.co"
	}

	slackWebHookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if slackWebHookURL == "" {
		fmt.Printf("SLACK_WEBHOOK_URL environment variable is not set\n")
		os.Exit(1)
	}

	hostname, err = os.Hostname()
	if err != nil {
		hostname = "<unknown hostname>"
	}
	fmt.Printf("The current machine's hostname is '%s'\n", hostname)

	if publicIPAddress, err = public.IPAddress(ifconfigURL); err != nil {
		fmt.Printf("client: could not get public IP address: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Public IP address is: '%s'\n", publicIPAddress)
	message := `The public IP address of ` + hostname + ` is: ` + publicIPAddress

	if err = slack.SendMessage(slackWebHookURL, hostname, message); err != nil {
		fmt.Printf("Failed to send slack a message: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("The message '%s' has been send to Slack OK\n", message)

	os.Exit(0)
}
