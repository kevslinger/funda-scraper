package database

import (
	"log"

	"github.com/jackc/pgx"
)

type Database struct {
	config    Config
	pgxConfig pgx.ConnConfig
}

func NewDatabase(config Config) Database {
	pgxConfig := pgx.ConnConfig{
		Host:     config.host,
		Port:     uint16(config.port),
		Database: config.name,
		User:     config.user,
		Password: config.password,
	}
	return Database{
		config:    config,
		pgxConfig: pgxConfig,
	}
}

func (d Database) SelectHouses() error {
	conn, err := d.connect()
	if err != nil {
		return err
	}
	rows, err := conn.Query("SELECT * FROM funda_houses")
	if err != nil {
		return err
	}
	log.Print("Rows: ", rows)
	for rows.Next() {
		row := rows.Scan()
		log.Print("Row: ", row)
	}
	return err
}

func (d Database) connect() (*pgx.Conn, error) {
	return pgx.Connect(d.pgxConfig)
}
