package main

import (
	"fmt"

	"github.com/SixteMorio/BotDiscord/bot"
	"github.com/SixteMorio/BotDiscord/config"
)

func main() {
	err := config.OctoPute()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	database.ConnectDb()
	bot.Start()
	
	<-make(chan struct{})

	return
}
