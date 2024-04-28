package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jasonlvhit/gocron"
	"github.com/kevslinger/funda-scraper/alerter"
	"github.com/kevslinger/funda-scraper/config"
	"github.com/kevslinger/funda-scraper/database"
	"github.com/kevslinger/funda-scraper/scraper"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags:  getFlags(),
		Action: daemon,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func daemon(ctx *cli.Context) error {
	config := config.LoadConfig(ctx)
	gocron.Every(config.GeneralConfig.ScrapeFrequency).Minutes().Do(scrapeFunda, config)
	<-gocron.Start()
	return nil
}

func scrapeFunda(config config.Config) error {

	fundaScraper := scraper.New(*config.ScraperConfig, &http.Client{})
	urls, err := fundaScraper.GetListingUrls(http.MethodGet, nil)
	if err != nil {
		return err
	}
	log.Print("Urls: ", urls)
	fundaListings := make([]scraper.FundaListing, 0)
	for _, url := range urls {
		fundaListing, err := fundaScraper.GetFundaListingFromUrl(url)
		if err != nil {
			log.Print("Error getting Funda Listing for house with URL ", url)
			continue
		}
		//log.Print("Funda listing created: ", fundaListing)
		fundaListings = append(fundaListings, fundaListing)
	}
	// TODO: Check with DB to see which listings need to be inserted/alerted

	alerts, err := alerter.New(alerter.Config{DiscordAuthenticationToken: config.AlerterConfig.DiscordAuthenticationToken, DiscordChannelID: config.AlerterConfig.DiscordChannelID})
	if err != nil {
		return err
	}
	for _, listing := range fundaListings[:config.GeneralConfig.NumHousesLimit] {
		alerts.Alert("Found a new house at " + listing.Address + ": " + listing.URL)
	}
	log.Print("Connecting to DB")
	db := database.NewDatabase(*config.DatabaseConfig)
	err = db.SelectHouses()
	if err != nil {
		log.Print("Error selecting houses: ", err)
	}
	err = db.InsertListings(fundaListings)
	if err != nil {
		log.Print("Error inserting listings: ", err)
	}
	return err
}

func getFlags() []cli.Flag {
	var flags []cli.Flag
	flags = append(scraper.Flags, config.Flags...)
	flags = append(flags, alerter.Flags...)
	flags = append(flags, database.Flags...)
	return flags
}
