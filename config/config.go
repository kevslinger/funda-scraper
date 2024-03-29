package config

import (
	"github.com/urfave/cli/v2"
)

var (
	configFileFlag = &cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

/*
We need configs for the following modules:

Alerter:
-

Database:
-

Scraper:
-

Main:
-


*/
