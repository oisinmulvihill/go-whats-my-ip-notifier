package settings

import (
	"os"
)

type configuration struct {
	IFConfigURL     string
	SlackWebHookURL string
	Hostname        string
}

func FromEnv() *configuration {
	ifconfigURL := os.Getenv("IFCONFIG_URL")
	if ifconfigURL == "" {
		ifconfigURL = "http://ifconfig.co"
	}

	slackWebHookURL := os.Getenv("SLACK_WEBHOOK_URL")

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "<unknown hostname>"
	}

	settings := configuration{
		IFConfigURL:     ifconfigURL,
		SlackWebHookURL: slackWebHookURL,
		Hostname:        hostname,
	}

	return &settings
}
