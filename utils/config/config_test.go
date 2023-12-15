package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/wlsyne/logIt/constants"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {

	f, err := os.CreateTemp("", "logItConfig.json")
	if err != nil {
		assert.FailNow(t, "Error creating temporary config file", err)
	}

	defer os.Remove(f.Name())

	//	when configFile not exist, this function should return nil and error
	_, err = GetConfig("notExist.json")
	assert.EqualError(t, err, "open notExist.json: no such file or directory")

	//	otherwise, it should return config as a map and nil
	mockData := `{
	  "Title": "My App",
	  "GitBaseUrl": "https://www.example.com",
	  "ChatIds": [123, 456],
	  "BotWebhook": "https://www.example.com/webhook"
	}`
	_, err = f.WriteString(mockData)
	if err != nil {
		assert.FailNow(t, "Error writing mock data to temporary config file", err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		assert.FailNow(t, "Error seeking to the beginning of file", err)
	}

	value, err := GetConfig(f.Name())
	assert.NoErrorf(t, err, "Expected nil, got error %v", err)

	//	value should be map[string]string
	assert.IsTypef(t, Config{}, value, "Expected map[string]string, got %v", value)

}

func TestValidator(t *testing.T) {
	//	title : string optional
	//	gitBaseUrl: string url required
	//	chatIds: when publish chatIds is required
	//	botWebhook: when publish botWebhook is required

	mockConfig := Config{
		GitBaseUrl: "https://www.abc.com",
		ChatIds:    []int{123, 456},
		BotWebhook: "https://www.abc.com",
	}

	//Check Title
	err := Validator(mockConfig, constants.WriteMode)
	assert.NoErrorf(t, err, "Expected nil, got error %v", err)

	//Check gitBaseUrl
	mockConfig.GitBaseUrl = "abc"
	err = Validator(mockConfig, constants.WriteMode)
	assert.EqualError(t, err, "GitBaseUrl is not a valid URL")

	//Check chatIds
	mockConfig.GitBaseUrl = "https://www.abc.com"
	mockConfig.ChatIds = []int{}

	// when Write
	err = Validator(mockConfig, constants.WriteMode)
	assert.NoErrorf(t, err, "Expected nil, got error %v", err)

	//When Publish chatIds is required
	err = Validator(mockConfig, constants.PublishMode)
	assert.EqualError(t, err, "ChatIds is required when publish")

	//Check botWebhook
	mockConfig.ChatIds = []int{123, 456}
	mockConfig.BotWebhook = "abc"

	//When Write
	err = Validator(mockConfig, constants.WriteMode)
	assert.NoErrorf(t, err, "Expected nil, got error %v", err)

	//	When Publish botWebhook is required
	err = Validator(mockConfig, constants.PublishMode)
	assert.EqualError(t, err, "BotWebhook is required when publish")
}
