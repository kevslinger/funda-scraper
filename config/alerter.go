package config

import (
	"github.com/urfave/cli/v2"
)

type AlerterConfig struct {
	WebexRoomId                string // TODO: Implementation
	DiscordAuthenticationToken string
	DiscordChannelID           string
	TelegramId                 string // TODO: Implementation
}

var AlerterDefaults *AlerterConfig = &AlerterConfig{}

const (
	discordAuthFlag      = "discord-auth-token"
	discordChannelIdFlag = "discord-channel-id"
)

var AlerterFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  discordAuthFlag,
		Usage: "The authentication token for a discord bot",
	},
	// TODO: Support multiple channels?
	&cli.StringFlag{
		Name:  discordChannelIdFlag,
		Usage: "The channel ID of the Discord channel to alert",
	},
}

func LoadAlerterConfig(ctx *cli.Context) *AlerterConfig {
	c := AlerterDefaults
	// Set ENV vars before command-line flags
	readStringFromEnv(discordAuthFlag, &c.DiscordAuthenticationToken)
	readStringFromEnv(discordChannelIdFlag, &c.DiscordChannelID)
	// Overwrite ENV vars on command-line
	if ctx.IsSet(discordAuthFlag) {
		c.DiscordAuthenticationToken = ctx.String(discordAuthFlag)
	}
	if ctx.IsSet(discordChannelIdFlag) {
		c.DiscordChannelID = ctx.String(discordChannelIdFlag)
	}
	return c
}
