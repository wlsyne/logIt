package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {

	f, err := os.CreateTemp("", "logItConfig.json")
	if err != nil {
		t.Fatal("Error creating temporary config file", err)
	}

	defer os.Remove(f.Name())

	//	when configFile not exist, this function should return nil and error
	_, err = GetConfig("notExist.json")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	//	otherwise, it should return config as a map and nil
	mockData := `{"title":"test","gitBaseUrl":"https://www.abc.com",chatIds:[123,456],botWebhook:"https://www.abc.com"}`
	_, err = f.WriteString(mockData)
	if err != nil {
		t.Fatal("Error writing mock data to temporary config file", err)
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		t.Fatal("Error seeking to the beginning of file", err)
	}

	value, err := GetConfig(f.Name())
	if err != nil {
		t.Error("Expected nil, got error", err)
	}
	//	value should be map[string]string
	_, ok := value.(map[string]interface{})
	if !ok {
		t.Error("Expected map[string]string, got", value)
	}

}

func TestCheckConfig(t *testing.T) {
	//	title : string optional
	//	gitBaseUrl: string url required
	//	chatIds: when publish chatIds is required
	//	botWebhook: when publish botWebhook is required
	mockConfig := map[string]interface{}{
		"gitBaseUrl": "https://www.abc.com",
		"chatIds":    []int{123, 456},
		"botWebhook": "https://www.abc.com",
	}

	//Check Title
	_, err := CheckConfig(mockConfig, Write)
	if err != nil {
		t.Error("Expected nil, got error", err)
	}

	//Check gitBaseUrl
	mockConfig["gitBaseUrl"] = "abc"
	_, err = CheckConfig(mockConfig, Write)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	//Check chatIds
	mockConfig["gitBaseUrl"] = "https://www.abc.com"
	mockConfig["chatIds"] = []int{}

	// when Write
	_, err = CheckConfig(mockConfig, Write)
	if err != nil {
		t.Error("Expected nil, got error", err)
	}

	//When Publish chatIds is required
	_, err = CheckConfig(mockConfig, Publish)
	if err == nil {
		t.Error("Expected error, got nil", err)
	}

	//Check botWebhook
	mockConfig["chatIds"] = []int{123, 456}
	mockConfig["botWebhook"] = "abc"

	//When Write
	_, err = CheckConfig(mockConfig, Write)
	if err != nil {
		t.Error("Expected nil, got error", err)
	}

	//	When Publish botWebhook is required
	_, err = CheckConfig(mockConfig, Publish)
	if err == nil {
		t.Error("Expected error, got nil", err)
	}
}
