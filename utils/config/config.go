package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wlsyne/logIt/constants"
	"os"
	"regexp"
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

	fmt.Println(config)

	return config, nil
}

func Validator(config Config, mode constants.Mode) error {
	//	title : string optional
	//	gitBaseUrl: string url required
	//	chatIds: when publish chatIds is required
	//	botWebhook: when publish botWebhook is required

	//Check gitBaseUrl
	match, err := regexp.MatchString(constants.RegexpStringUrl, config.GitBaseUrl)
	if err != nil || !match {
		return errors.New("GitBaseUrl is not a valid URL")
	}

	//Check chatIds
	if mode == constants.PublishMode && len(config.ChatIds) == 0 {
		return errors.New("ChatIds is required when publish")
	}

	//Check botWebhook
	match, err = regexp.MatchString(constants.RegexpStringUrl, config.BotWebhook)
	if (err != nil || !match) && mode == constants.PublishMode {
		return errors.New("BotWebhook is required when publish")
	}

	return nil
}
