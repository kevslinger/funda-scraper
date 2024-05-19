package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/kevslinger/funda-scraper/alerter"
	"github.com/kevslinger/funda-scraper/config"
	"github.com/kevslinger/funda-scraper/database"
	"github.com/kevslinger/funda-scraper/scraper"
	"github.com/urfave/cli/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "err", err)
	}
	app := &cli.App{
		Flags:  getFlags(),
		Action: daemon,
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error(err.Error())
		return
	}
}

func daemon(ctx *cli.Context) error {
	c := config.LoadConfig(ctx)
	fundaScraper := scraper.New(*c.ScraperConfig, &http.Client{})
	db := database.New(*c.DatabaseConfig)
	alerts, err := alerter.New(*c.AlerterConfig)
	if err != nil {
		return err
	}

	slog.Info("Starting funda-scraper! New houses will be scraped every", "minutes", c.GeneralConfig.ScrapeFrequency)
	gocron.Every(uint64(c.GeneralConfig.ScrapeFrequency)).Minutes().Do(scrapeFunda, fundaScraper, db, alerts, c)
	<-gocron.Start()
	return nil
}

func scrapeFunda(fundaScraper *scraper.Scraper, db database.Database, alerts *alerter.Alerter, config config.Config) error {
	urls, err := fundaScraper.GetListingUrls(http.MethodGet, nil)
	if err != nil {
		slog.Error("Error getting listing URLs", "err", err)
		return err
	}
	slog.Info("Got listing URLs", "URLs", urls)
	fundaListings := getNewFundaListings(fundaScraper, db, urls)
	if len(fundaListings) == 0 {
		slog.Warn("No new houses")
		return nil
	} else {
		slog.Info("Found new houses", "num", len(fundaListings))
	}

	rowsInserted, err := db.InsertListings(fundaListings)
	if err != nil {
		slog.Warn("Error inserting listings", "err", err)
	} else {
		slog.Info("Committed new listings to DB", "numNewListings", rowsInserted)
	}

	for _, listing := range fundaListings[:config.GeneralConfig.NumHousesLimit] {
		alerts.Alert("Found a new house at " + listing.Address + ": " + listing.URL)
	}

	return err
}

func getFlags() []cli.Flag {
	var flags []cli.Flag
	flags = append(config.ScraperFlags, config.GeneralFlags...)
	flags = append(flags, config.AlerterFlags...)
	flags = append(flags, config.DatabaseFlags...)
	return flags
}

func getNewFundaListings(fundaScraper *scraper.Scraper, db database.Database, urls []string) []scraper.FundaListing {
	var fundaListings []scraper.FundaListing
	for _, url := range urls {
		fundaListing, err := fundaScraper.GetFundaListingFromUrl(url)
		if err != nil {
			slog.Warn("Error getting funda listing for house", "url", url, "err", err)
			continue
		}
		// TODO: Have some sort of time limit?
		alreadyPresent, err := db.SelectHouseWithAddress(fundaListing.Address)
		if err != nil {
			slog.Warn("Error selecting house with address", "err", err)
		}
		if alreadyPresent == "" {
			fundaListings = append(fundaListings, fundaListing)
		}
	}
	return fundaListings
}
