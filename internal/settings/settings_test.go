package settings

import (
	"testing"

	"github.com/oisinmulvihill/go-whats-my-ip-notifier/internal/assert"
)

func TestGetSettingFromEnvDefaultValues(t *testing.T) {
	config := FromEnv()
	assert.Equal(t, config.IFConfigURL, "http://ifconfig.co")
	assert.Equal(t, config.SlackWebHookURL, "")
}

func TestGetSettingFromEnvCustomValues(t *testing.T) {
	t.Setenv("IFCONFIG_URL", "http://example.com")
	t.Setenv("SLACK_WEBHOOK_URL", "http://example.com/slack")

	config := FromEnv()
	assert.Equal(t, config.IFConfigURL, "http://example.com")
	assert.Equal(t, config.SlackWebHookURL, "http://example.com/slack")
}
