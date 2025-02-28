package settings

import (
	"os"
)

type configuration struct {
	IFConfigURL     string
	SlackWebHookURL string
	Hostname        string
	StorageFilePath string
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

	storageFilePath := os.Getenv("STORAGE_FILE_PATH")
	if storageFilePath == "" {
		storageFilePath = "storage.db"
	}

	settings := configuration{
		IFConfigURL:     ifconfigURL,
		SlackWebHookURL: slackWebHookURL,
		Hostname:        hostname,
		StorageFilePath: storageFilePath,
	}

	return &settings
}
