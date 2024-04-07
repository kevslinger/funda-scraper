package config

import (
	"github.com/kevslinger/funda-scraper/alerter"
	"github.com/kevslinger/funda-scraper/database"
	"github.com/kevslinger/funda-scraper/scraper"
	"github.com/naoina/toml"
	"github.com/urfave/cli/v2"
)

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	},
	&cli.IntFlag{
		Name:  "num-houses-limit",
		Usage: "The number of new houses to receive on each update",
		Value: 10,
	},
}

type GeneralConfig struct {
	TomlConfigPath string
	NumHousesLimit int
}

var Defaults *GeneralConfig = &GeneralConfig{}

var tomlSettings = toml.Config{}

type Config struct {
	AlerterConfig  *alerter.Config
	DatabaseConfig *database.Config
	ScraperConfig  *scraper.Config
	GeneralConfig  *GeneralConfig
}

func LoadConfig(ctx *cli.Context) Config {
	return Config{
		AlerterConfig:  alerter.LoadConfig(ctx),
		DatabaseConfig: database.LoadConfig(ctx),
		ScraperConfig:  scraper.LoadConfig(ctx),
		GeneralConfig:  loadConfig(ctx),
	}
}

func loadConfig(ctx *cli.Context) *GeneralConfig {
	config := Defaults
	if ctx.IsSet("config") {
		config.TomlConfigPath = ctx.String("config")
	}
	if ctx.IsSet("num-houses-limit") {
		config.NumHousesLimit = ctx.Int("num-houses-limit")
	}
	return config
}
