package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Title      string
	GitBaseUrl string
	ChatIds    []int
	BotWebhook string
}

func GetConfig(configFilePath string) (Config, error) {
	// Read File
	info, err := os.ReadFile(configFilePath)

	if err != nil {
		return Config{}, err
	}

	//Convert info to Config and info is a JSON string
	var config Config
	err = json.Unmarshal(info, &config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
