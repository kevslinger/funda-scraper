package config

import (
	"github.com/kevslinger/funda-scraper/alerter"
	"github.com/kevslinger/funda-scraper/database"
	"github.com/kevslinger/funda-scraper/scraper"
	"github.com/naoina/toml"
	"github.com/urfave/cli/v2"
)

var (
	configFileFlag = &cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

var tomlSettings = toml.Config{}

type Config struct {
	AlerterConfig  *alerter.Config
	DatabaseConfig *database.Config
	ScraperConfig  *scraper.Config
}

func LoadConfig(ctx *cli.Context) Config {
	return Config{
		AlerterConfig:  alerter.LoadConfig(ctx),
		DatabaseConfig: database.LoadConfig(ctx),
		ScraperConfig:  scraper.LoadConfig(ctx),
	}
}
