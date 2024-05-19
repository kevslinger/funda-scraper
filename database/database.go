package database

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kevslinger/funda-scraper/config"
	"github.com/kevslinger/funda-scraper/scraper"
)

type Database struct {
	config config.DatabaseConfig
}

func New(config config.DatabaseConfig) Database {
	return Database{
		config: config,
	}
}

func (d Database) SelectHouses() error {
	conn, err := d.connect()
	if err != nil {
		return err
	}
	defer conn.Close(context.TODO())

	listing := scraper.FundaListing{}
	err = conn.QueryRow(context.TODO(), "SELECT link, house_address, house_description FROM funda_houses").Scan(&listing.URL, &listing.Address, &listing.Description)
	if err != nil {
		return err
	}
	slog.Info("TEST from DB got ", "listing", listing)
	return err
}

func (d Database) SelectHouseWithLink(link string) (string, error) {
	conn, err := d.connect()
	if err != nil {
		return "", err
	}
	defer conn.Close(context.TODO())
	var foundLink string
	err = conn.QueryRow(context.TODO(), "SELECT link from funda_houses WHERE link=$1", link).Scan(&foundLink)
	return foundLink, err
}

func (d Database) SelectHouseWithAddress(address string) (string, error) {
	conn, err := d.connect()
	if err != nil {
		return "", err
	}
	defer conn.Close(context.TODO())

	var foundAddress string
	err = conn.QueryRow(context.TODO(), "SELECT house_address from funda_houses WHERE house_address=$1", address).Scan(&foundAddress)
	return foundAddress, err
}

func (d Database) InsertListings(listings []scraper.FundaListing) (int, error) {
	conn, err := d.connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close(context.TODO())

	copyCount, err := conn.CopyFrom(
		context.TODO(),
		pgx.Identifier{"funda_houses"},
		[]string{"time_seen", "link", "house_address", "price", "house_description", "zip_code", "built_year", "total_size", "living_size", "house_type", "building_type", "num_rooms", "num_bedrooms"},
		pgx.CopyFromSlice(len(listings), func(i int) ([]any, error) {
			return []any{
				time.Now().UTC(),
				escapeStringForQuery(listings[i].URL),
				escapeStringForQuery(listings[i].Address),
				listings[i].Price,
				escapeStringForQuery(listings[i].Description),
				escapeStringForQuery(listings[i].ZipCode),
				listings[i].BuildYear,
				listings[i].TotalSize,
				listings[i].LivingSize,
				escapeStringForQuery(listings[i].HouseType),
				escapeStringForQuery(listings[i].BuildingType),
				listings[i].NumRooms,
				listings[i].NumBedrooms,
			}, nil
		}),
	)

	return int(copyCount), err
}

func (d Database) connect() (*pgx.Conn, error) {
	return pgx.Connect(context.TODO(), getDatabaseURL(d.config.User, d.config.Password, d.config.Host, d.config.Name, d.config.Port))
}

func getDatabaseURL(user, password, host, name string, port int) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, name)
}

func escapeStringForQuery(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
