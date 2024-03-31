package scraper

import "github.com/urfave/cli/v2"

type Config struct {
	base_url string
	headers  map[string]string
}

var Defaults *Config = &Config{
	base_url: "https://www.funda.nl",
	headers:  map[string]string{"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"},
}

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  "listing-path",
		Usage: "The path of the listing to search in Funda",
	},
}
