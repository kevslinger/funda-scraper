package config

import (
	"github.com/urfave/cli/v2"
)

const (
	hostFlag     = "db-host"
	portFlag     = "db-port"
	nameFlag     = "db-name"
	userFlag     = "db-user"
	passwordFlag = "db-password"
)

var DatabaseFlags []cli.Flag = []cli.Flag{
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

type DatabaseConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

var DatabaseDefaults *DatabaseConfig = &DatabaseConfig{}

func LoadDatabaseConfig(ctx *cli.Context) *DatabaseConfig {
	c := DatabaseDefaults
	// Set ENV vars before command-line flags
	readStringFromEnv(hostFlag, &c.Host)
	readIntFromEnv(portFlag, &c.Port)
	readStringFromEnv(nameFlag, &c.Name)
	readStringFromEnv(userFlag, &c.User)
	readStringFromEnv(passwordFlag, &c.Password)
	// Overwrite ENV vars on command-line
	if ctx.IsSet(hostFlag) {
		c.Host = ctx.String(hostFlag)
	}
	if ctx.IsSet(portFlag) {
		c.Port = ctx.Int(portFlag)
	}
	if ctx.IsSet(nameFlag) {
		c.Name = ctx.String(nameFlag)
	}
	if ctx.IsSet(userFlag) {
		c.User = ctx.String(userFlag)
	}
	if ctx.IsSet(passwordFlag) {
		c.Password = ctx.String(passwordFlag)
	}
	return c
}
