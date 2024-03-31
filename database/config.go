package database

type Config struct {
	host     string
	port     int
	name     string
	user     string
	password string
}

var Defaults *Config = &Config{}
