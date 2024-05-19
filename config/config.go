package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	numHousesFlag       = "num-houses-limit"
	scrapeFrequencyFlag = "scrape-frequency-minutes"
)

var GeneralFlags []cli.Flag = []cli.Flag{
	&cli.IntFlag{
		Name:  numHousesFlag,
		Usage: "The number of new houses to receive on each update",
		Value: 10,
	},
	&cli.IntFlag{
		Name:  scrapeFrequencyFlag,
		Usage: "The number of minutes in between scheduling the scraper and alerter to run",
		Value: 60,
	},
}

type GeneralConfig struct {
	NumHousesLimit  int
	ScrapeFrequency int
}

var Defaults *GeneralConfig = &GeneralConfig{}

type Config struct {
	AlerterConfig  *AlerterConfig
	DatabaseConfig *DatabaseConfig
	ScraperConfig  *ScraperConfig
	GeneralConfig  *GeneralConfig
}

func LoadConfig(ctx *cli.Context) Config {
	return Config{
		AlerterConfig:  LoadAlerterConfig(ctx),
		DatabaseConfig: LoadDatabaseConfig(ctx),
		ScraperConfig:  LoadScraperConfig(ctx),
		GeneralConfig:  loadConfig(ctx),
	}
}

func loadConfig(ctx *cli.Context) *GeneralConfig {
	c := Defaults
	// Use ENV vars which get overriden if the flags are present
	readIntFromEnv(numHousesFlag, &c.NumHousesLimit)
	readIntFromEnv(scrapeFrequencyFlag, &c.ScrapeFrequency)
	// Overwrite ENV vars on command-line
	if ctx.IsSet(numHousesFlag) {
		c.NumHousesLimit = ctx.Int(numHousesFlag)
	}
	if ctx.IsSet(scrapeFrequencyFlag) {
		c.ScrapeFrequency = ctx.Int(scrapeFrequencyFlag)
	}
	return c
}

// readIntFromEnv reads an ENV var into an integer, if possible
func readIntFromEnv(flagName string, configPtr *int) {
	if envVal, ok := os.LookupEnv(convertFlagNameToEnvVar(flagName)); ok {
		if envValInt, err := strconv.Atoi(envVal); err == nil {
			*configPtr = envValInt
		} else {
			slog.Warn("Error reading env var", "name", flagName, "value", envVal, "err", err)
		}
	}
}

// readStringFromEnv reads an ENV var into a string
func readStringFromEnv(flagName string, configPtr *string) {
	if envVal, ok := os.LookupEnv(convertFlagNameToEnvVar(flagName)); ok {
		*configPtr = envVal
	}
}

// readStringSliceFromEnv reads a comma-separated ENV var into a string slice
func readStringSliceFromEnv(flagName string, configPtr *[]string) {
	if envVal, ok := os.LookupEnv(convertFlagNameToEnvVar(flagName)); ok {
		*configPtr = strings.Split(envVal, ",")
	}
}

// convertFlagNameToEnvVar changes flag names (e.g. "max-price") to env names (e.g. "MAX_PRICE")
func convertFlagNameToEnvVar(flagName string) string {
	return strings.ToUpper(strings.ReplaceAll(flagName, "-", "_"))
}
