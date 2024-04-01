package alerter

import "github.com/urfave/cli/v2"

type Config struct {
	WebexRoomId                string // TODO: Implementation
	DiscordAuthenticationToken string
	DiscordChannelID           string
	TelegramId                 string // TODO: Implementation
}

var Defaults *Config = &Config{}

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  "discord-auth-token",
		Usage: "The authentication token for a discord bot",
	},
	// TODO: Support multiple channels?
	&cli.StringFlag{
		Name:  "discord-channel-id",
		Usage: "The channel ID of the Discord channel to alert",
	},
}

// TODO: Better way to set Config vars from context?
func LoadConfig(ctx *cli.Context) *Config {
	config := Defaults
	if ctx.IsSet("discord-auth-token") {
		config.DiscordAuthenticationToken = ctx.String("discord-auth-token")
	}
	if ctx.IsSet("discord-channel-id") {
		config.DiscordChannelID = ctx.String("discord-channel-id")
	}
	return config
}
