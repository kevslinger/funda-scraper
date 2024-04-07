package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kevslinger/funda-scraper/alerter"
	"github.com/kevslinger/funda-scraper/config"
	"github.com/kevslinger/funda-scraper/scraper"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags:  getFlags(),
		Action: scrapeFunda,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func scrapeFunda(ctx *cli.Context) error {
	config := config.LoadConfig(ctx)
	fundaScraper := scraper.New(*config.ScraperConfig, &http.Client{})
	urls, err := fundaScraper.GetListingUrls(http.MethodGet, nil)
	if err != nil {
		return err
	}
	alerts, err := alerter.New(alerter.Config{DiscordAuthenticationToken: ctx.String("discord-auth-token"), DiscordChannelID: ctx.String("discord-channel-id")})
	if err != nil {
		return err
	}
	for _, url := range urls[:config.GeneralConfig.NumHousesLimit] {
		alerts.Alert("Found a new house at " + url)
	}
	// fundaListings, err := fundaScraper.ParseResponse(response)
	// if err != nil {
	// 	return err
	// }
	// for _, listing := range fundaListings {
	// 	alerts.Alert("Found a new house at " + listing.URL)
	// }
	return nil
}

func getFlags() []cli.Flag {
	var flags []cli.Flag
	flags = append(scraper.Flags, config.Flags...)
	flags = append(flags, alerter.Flags...)
	return flags
}
