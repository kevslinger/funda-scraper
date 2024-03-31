package alerter

type Alerter struct {
	discord          *DiscordBot
	discordChannelID string
}

func New(config Config) (*Alerter, error) {
	discord, err := NewDiscordBot(config.DiscordAuthenticationToken)
	if err != nil {
		return nil, err
	}
	return &Alerter{discord: discord, discordChannelID: config.DiscordChannelID}, nil
}

func (a Alerter) Alert(alert string) error {
	// Send alert to all the platforms we have
	if a.discord != nil && a.discordChannelID != "" {
		return a.discord.SendMessage(a.discordChannelID, alert)
	}
	return nil
}
