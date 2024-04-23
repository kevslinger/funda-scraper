package alerter

import "github.com/urfave/cli/v2"

type Config struct {
	WebexRoomId                string // TODO: Implementation
	DiscordAuthenticationToken string
	DiscordChannelID           string
	TelegramId                 string // TODO: Implementation
}

var Defaults *Config = &Config{}

const (
	discordAuthFlag  = "discord-auth-token"
	discordChannelID = "discord-channel-id"
)

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  discordAuthFlag,
		Usage: "The authentication token for a discord bot",
	},
	// TODO: Support multiple channels?
	&cli.StringFlag{
		Name:  discordChannelID,
		Usage: "The channel ID of the Discord channel to alert",
	},
}

// TODO: Better way to set Config vars from context?
func LoadConfig(ctx *cli.Context) *Config {
	config := Defaults
	if ctx.IsSet(discordAuthFlag) {
		config.DiscordAuthenticationToken = ctx.String(discordAuthFlag)
	}
	if ctx.IsSet(discordChannelID) {
		config.DiscordChannelID = ctx.String(discordChannelID)
	}
	return config
}
