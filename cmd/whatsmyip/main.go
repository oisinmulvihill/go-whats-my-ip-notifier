package main

import (
	"fmt"
	"log"
	"os"

	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/public"
	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/settings"
	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/slack"
	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/storage"
)

func main() {
	var err error
	var publicIPAddress string

	config := settings.FromEnv()
	if config.SlackWebHookURL == "" {
		log.Fatalln("SLACK_WEBHOOK_URL environment variable is not set")
		os.Exit(1)
	}

	// Create the ip logging db if it doesn't exist:
	ipStore, err := storage.Init(config.StorageFilePath)
	if err != nil {
		log.Fatalf("Could not create the ip logging db: %s\n", err)
		os.Exit(1)
	}

	log.Printf("The current machine's hostname is '%s'\n", config.Hostname)

	if publicIPAddress, err = public.IPAddress(config.IFConfigURL); err != nil {
		log.Fatalf("Could not get public IP address: %s\n", err)
		os.Exit(1)
	}
	log.Printf("The current public IP address is: '%s'\n", publicIPAddress)

	address, err := ipStore.CurrentIP()
	if err != nil {
		log.Fatalf("Could not get recover any stored IP address: %s\n", err)
		os.Exit(1)
	}
	log.Printf("The current stored IP address is: '%s'\n", address)

	if address != publicIPAddress {
		message := fmt.Sprintf("The IP address has changed from '%s' to '%s'\n", address, publicIPAddress)
		ipStore.AddAddress(publicIPAddress)
		if err = slack.SendMessage(config.SlackWebHookURL, config.Hostname, message); err != nil {
			log.Fatalf("Failed to send slack a message: %s\n", err)
			os.Exit(1)
		}
		log.Printf("%s", message)

	} else {
		log.Printf("The IP address has not changed from %s\n", address)
	}

	os.Exit(0)
}
