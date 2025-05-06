package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *Config
)

type Config struct {
	TOKEN      string `json:"TOKEN"`
	Bot_PREFIX string `json:"BOT_PREFIX"`
}

func OctoPute() error {
	fmt.Println("Loading config file")
	file, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println("Error while loading the file")
		return err
	}

	fmt.Println("Converting file to config struct")
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("Error while unmarshalling the file")
		return err
	}

	Token = config.TOKEN
	BotPrefix = config.Bot_PREFIX

	return nil
}
