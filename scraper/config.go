package scraper

import "github.com/urfave/cli/v2"

type Config struct {
	baseUrl string
	headers map[string]string
}

const baseUrl = "https://www.funda.nl"

var Defaults *Config = &Config{
	baseUrl: baseUrl,
	headers: map[string]string{"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"},
}

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  "base-url",
		Value: baseUrl,
		Usage: "The base URL to scan",
	},
	&cli.StringFlag{
		Name:  "listing-path",
		Usage: "The path of the listing to search in Funda",
	},
}

// TODO: Better way to load config from flags?
func LoadConfig(ctx *cli.Context) *Config {
	config := Defaults
	if ctx.IsSet("base-url") {
		config.baseUrl = ctx.String("base-url")
	}
	return config
}
