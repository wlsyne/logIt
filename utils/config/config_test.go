package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {

	f, err := os.CreateTemp("", "logItConfig.json")
	if err != nil {
		assert.FailNow(t, "Error creating temporary config file", err)
	}

	//defer os.Remove(f.Name())

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

//func TestCheckConfig(t *testing.T) {
//	//	title : string optional
//	//	gitBaseUrl: string url required
//	//	chatIds: when publish chatIds is required
//	//	botWebhook: when publish botWebhook is required
//	mockConfig := map[string]interface{}{
//		"gitBaseUrl": "https://www.abc.com",
//		"chatIds":    []int{123, 456},
//		"botWebhook": "https://www.abc.com",
//	}
//
//	//Check Title
//	_, err := CheckConfig(mockConfig, Write)
//	assert.NoErrorf(t, err, "Expected nil, got error %v", err)
//
//	//Check gitBaseUrl
//	mockConfig["gitBaseUrl"] = "abc"
//	_, err = CheckConfig(mockConfig, Write)
//	assert.EqualError(t, err, "gitBaseUrl is not a valid url")
//
//	//Check chatIds
//	mockConfig["gitBaseUrl"] = "https://www.abc.com"
//	mockConfig["chatIds"] = []int{}
//
//	// when Write
//	_, err = CheckConfig(mockConfig, Write)
//	assert.NoErrorf(t, err, "Expected nil, got error %v", err)
//
//	//When Publish chatIds is required
//	_, err = CheckConfig(mockConfig, Publish)
//	assert.EqualError(t, err, "chatIds is required when publish")
//
//	//Check botWebhook
//	mockConfig["chatIds"] = []int{123, 456}
//	mockConfig["botWebhook"] = "abc"
//
//	//When Write
//	_, err = CheckConfig(mockConfig, Write)
//	assert.NoErrorf(t, err, "Expected nil, got error %v", err)
//
//	//	When Publish botWebhook is required
//	_, err = CheckConfig(mockConfig, Publish)
//	assert.EqualError(t, err, "botWebhook is required when publish")
//}
