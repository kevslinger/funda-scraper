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
