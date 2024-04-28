package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/kevslinger/funda-scraper/scraper"
)

type Database struct {
	config Config
}

func NewDatabase(config Config) Database {
	return Database{
		config: config,
	}
}

func (d Database) SelectHouses() error {
	conn, err := d.connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	listing := scraper.FundaListing{}
	err = conn.QueryRow(context.Background(), "SELECT link, house_address, house_description FROM funda_houses").Scan(&listing.URL, &listing.Address, &listing.Description)
	if err != nil {
		return err
	}
	log.Print(listing)
	return err
}

func (d Database) selectHouseWithLink(link string, conn *pgx.Conn) (string, error) {
	var foundLink string
	err := conn.QueryRow(context.Background(), "SELECT link from funda_houses WHERE link=$1", link).Scan(&foundLink)
	return foundLink, err
}

func (d Database) InsertListings(listings []scraper.FundaListing) error {
	// TODO: Try CopyFrom https://github.com/jackc/pgx/blob/v5.5.5/copy_from.go#L265
	conn, err := d.connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	for _, listing := range listings {
		foundLink, _ := d.selectHouseWithLink(listing.URL, conn)
		if foundLink == listing.URL {
			continue
		}
		tx, err := conn.Begin(context.Background())
		if err != nil {
			return err
		}
		defer tx.Rollback(context.Background())

		// TODO: Is this weak against SQL injection?
		_, err = tx.Exec(context.Background(), fmt.Sprintf("insert into funda_houses(link, house_address, house_description) values ('%s', '%s', '%s')", escapeStringForQuery(listing.URL), escapeStringForQuery(listing.Address), escapeStringForQuery(listing.Description)))
		if err != nil {
			return err
		}

		err = tx.Commit(context.Background())
		if err != nil {
			return err
		}
		log.Print("Committed house with Address: ", listing.Address)
	}
	return nil
}

func (d Database) connect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), getDatabaseURL(d.config.user, d.config.password, d.config.host, d.config.name, d.config.port))
}

func getDatabaseURL(user, password, host, name string, port int) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, name)
}

func escapeStringForQuery(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
