package alerter

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordBot struct {
	bot *discordgo.Session
}

func NewDiscordBot(authToken string) (*DiscordBot, error) {
	// TODO: Is this expected behavior?
	if authToken == "" {
		return nil, nil
	}
	bot, err := discordgo.New("Bot " + authToken)
	if err != nil {
		return &DiscordBot{}, fmt.Errorf("error creating discord bot: %e", err)
	}
	return &DiscordBot{bot: bot}, nil
}

// TODO: Must cut down long strings into smaller bits
func (d DiscordBot) SendMessage(channelID, message string) error {
	err := d.bot.Open()
	if err != nil {
		return err
	}
	defer d.bot.Close()
	_, err = d.bot.ChannelMessageSend(channelID, message)
	return err
}
