#!/bin/bash
WHATSMYIP_EXE="whats-my-ip"
export SLACK_WEBHOOK_URL="https://hooks.slack.com/services/XXXXXXXXX/XXXXXXXXX/XXXXXXXXXXXXXXXXXXXXXXXX"
export STORAGE_FILE_PATH="/tmp/whatsmyip.db"

$WHATSMYIP_EXE
