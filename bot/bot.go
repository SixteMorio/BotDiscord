package bot

import (
	"fmt"
	// "time"

	"github.com/SixteMorio/BotDiscord/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "cc" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ocotputeuuuh !")
	}

	if m.Author.ID == "527083932311748618" && m.Content == config.BotPrefix + "insulte" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "EHH TOI LA QU'EST CE QUE TU FAIS LA , RESTE ICI !!")
		_, _ = s.ChannelMessageSend(m.ChannelID, "ENLEVE TON FUT, TU VAS TE PRENDRE UNE TURLUTE (rime + 10)\nhttps://tenor.com/view/shrek-smirk-secret-reversed-backwards-gif-17744525497828225476")
	}

	// Vérification si c'est le bot de Luca
	/*if m.Author.ID == "1369279358203854899" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "EH OH ! LE BOT DE LUCA !")
		time.Sleep(1 * time.Second)
		_, _ = s.ChannelMessageSend(m.ChannelID, "LUCA, TU VAS TE FAIRE VIRER !")
		time.Sleep(1 * time.Second)
		err := s.GuildMemberRemove(m.GuildID, "527083932311748618", "Bot de Luca détecté")
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Erreur lors de la suppression de Luca: " + err.Error())
			fmt.Println("Erreur de suppression:", err)
		}
		return
	}*/

	// Phase de pressing pour chaque message
	/*_, _ = s.ChannelMessageSend(m.ChannelID, "EH OH !")
	time.Sleep(1 * time.Second)
	_, _ = s.ChannelMessageSend(m.ChannelID, "QU'EST-CE QUE TU DIS ?")
	time.Sleep(1 * time.Second)
	_, _ = s.ChannelMessageSend(m.ChannelID, "TU VAS ME LE RÉPÉTER ?")
	time.Sleep(1 * time.Second)
	_, _ = s.ChannelMessageSend(m.ChannelID, "https://giphy.com/gifs/j7dsfGlRnzq4EHZjhK")*/
}
