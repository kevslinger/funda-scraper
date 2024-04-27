package database

import "github.com/urfave/cli/v2"

const (
	hostFlag     = "db-host"
	portFlag     = "db-port"
	nameFlag     = "db-name"
	userFlag     = "db-user"
	passwordFlag = "db-password"
)

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  hostFlag,
		Usage: "IP of the DB Host",
		Value: "localhost",
	},
	&cli.IntFlag{
		Name:  portFlag,
		Usage: "Port to connect to the DB",
	},
	&cli.StringFlag{
		Name:  nameFlag,
		Usage: "Name of the DB",
	},
	&cli.StringFlag{
		Name:  userFlag,
		Usage: "DB Username",
	},
	&cli.StringFlag{
		Name:  passwordFlag,
		Usage: "DB Password",
	},
}

type Config struct {
	host     string
	port     int
	name     string
	user     string
	password string
}

var Defaults *Config = &Config{}

func LoadConfig(ctx *cli.Context) *Config {
	config := Defaults
	if ctx.IsSet(hostFlag) {
		config.host = ctx.String(hostFlag)
	}
	if ctx.IsSet(portFlag) {
		config.port = ctx.Int(portFlag)
	}
	if ctx.IsSet(nameFlag) {
		config.name = ctx.String(nameFlag)
	}
	if ctx.IsSet(userFlag) {
		config.user = ctx.String(userFlag)
	}
	if ctx.IsSet(passwordFlag) {
		config.password = ctx.String(passwordFlag)
	}
	return config
}
