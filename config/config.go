package config

import (
	"github.com/kevslinger/funda-scraper/alerter"
	"github.com/kevslinger/funda-scraper/database"
	"github.com/kevslinger/funda-scraper/scraper"
	"github.com/naoina/toml"
	"github.com/urfave/cli/v2"
)

const (
	configFlag          = "config"
	numHousesFlag       = "num-houses-limit"
	scrapeFrequencyFlag = "scrape-frequency-minutes"
)

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  configFlag,
		Usage: "TOML configuration file",
	},
	&cli.IntFlag{
		Name:  numHousesFlag,
		Usage: "The number of new houses to receive on each update",
		Value: 10,
	},
	&cli.Uint64Flag{
		Name:  scrapeFrequencyFlag,
		Usage: "The number of minutes in between scheduling the scraper and alerter to run",
		Value: 60,
	},
}

type GeneralConfig struct {
	TomlConfigPath  string
	NumHousesLimit  int
	ScrapeFrequency uint64
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
	if ctx.IsSet(configFlag) {
		config.TomlConfigPath = ctx.String(configFlag)
	}
	if ctx.IsSet(numHousesFlag) {
		config.NumHousesLimit = ctx.Int(numHousesFlag)
	}
	if ctx.IsSet(scrapeFrequencyFlag) {
		config.ScrapeFrequency = ctx.Uint64(scrapeFrequencyFlag)
	}
	return config
}
