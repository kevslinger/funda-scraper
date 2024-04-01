package database

import "github.com/urfave/cli/v2"

type Config struct {
	host     string
	port     int
	name     string
	user     string
	password string
}

var Defaults *Config = &Config{}

func LoadConfig(ctx *cli.Context) *Config {
	return Defaults
}
