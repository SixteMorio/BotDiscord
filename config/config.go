package config

import (
	"os"
)

var (
	Token     string
	BotPrefix string

)

func OctoPute() error {
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	return nil
}
