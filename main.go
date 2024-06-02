package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
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

	slog.Info("Running funda-scraper!")
	scrapeFunda(fundaScraper, db, alerts, c)
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
	newListings, listingsToReplace := getNewFundaListings(fundaScraper, db, urls, config.GeneralConfig.HouseLookbackDays)
	allListings := append(newListings, listingsToReplace...)
	slog.Info("postcodes selected", "codes", config.ScraperConfig.Postcode)
	if len(allListings) == 0 {
		slog.Warn("No new houses")
		return nil
	} else {
		slog.Info("Found new houses", "numNew", len(newListings), "numToUpdate", len(listingsToReplace))
	}

	rowsInserted, err := db.InsertListings(newListings)
	if err != nil {
		slog.Warn("Error inserting listings", "err", err)
	} else {
		slog.Info("Committed new listings to DB", "numNewListings", rowsInserted)
	}
	rowsUpdated, err := db.UpdateListings(listingsToReplace)
	if err != nil {
		slog.Warn("Error updating listings", "err", err)
	} else {
		slog.Info("Updating listings in DB", "numUpdatedListings", rowsUpdated)
	}

	var index, numAlerted int
	for index < len(allListings) && (numAlerted < config.GeneralConfig.NumHousesLimit || config.GeneralConfig.NumHousesLimit < 0) {
		var shouldAlert bool
		// Only alert houses that are in the postcodes listed, if any
		if len(config.ScraperConfig.Postcode) > 0 {
			listingPostcode := allListings[index].Postcode
			for _, postcode := range config.ScraperConfig.Postcode {
				if strings.Contains(listingPostcode, postcode) {
					shouldAlert = true
					break
				}
			}
		} else {
			shouldAlert = true
		}
		if shouldAlert {
			alerts.Alert("Found a new house at " + allListings[index].Address + ": " + allListings[index].URL)
			numAlerted++
		}
		index++
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

// getNewFundaListings takes a slice of funda URLs and returns a slice of funda listings. If lookbackDays is provided
// it will separate between listings that haven't been seen before, and those that have been (but longer than lookbackDays ago)
func getNewFundaListings(fundaScraper *scraper.Scraper, db database.Database, urls []string, lookbackDays int) ([]scraper.FundaListing, []scraper.FundaListing) {
	var listingsToReplace, newListings []scraper.FundaListing
	// TODO: Should I batch these queries?
	for _, url := range urls {
		fundaListing, err := fundaScraper.GetFundaListingFromUrl(url)
		if err != nil {
			slog.Warn("Error getting funda listing for house", "url", url, "err", err)
			continue
		}
		listingFoundInDb, err := db.SelectHouseWithAddress(fundaListing.Address)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			slog.Warn("Error selecting house with address", "addr", fundaListing.Address, "err", err)
			continue
		}
		// Lookback days turned off -- only care if we found the house at some point
		if lookbackDays < 0 {
			if errors.Is(err, pgx.ErrNoRows) {
				newListings = append(newListings, fundaListing)
			}
		} else {
			listingFoundRecently, err := db.SelectRecentHouseWithAddress(fundaListing.Address, lookbackDays)
			if err != nil && !errors.Is(err, pgx.ErrNoRows) {
				slog.Warn("Error selecting house with address", "addr", fundaListing.Address, "err", err)
				continue
			}
			// Case 1: Not found in db at all
			if listingFoundRecently == "" && listingFoundInDb == "" {
				newListings = append(newListings, fundaListing)
				// Case 2: Found in DB but not recently
			} else if listingFoundRecently == "" && listingFoundInDb != "" {
				listingsToReplace = append(listingsToReplace, fundaListing)
			}
			// Case 3: Found in DB recently -- no action required
		}
	}
	return newListings, listingsToReplace
}
