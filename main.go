package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/kevslinger/funda-scraper/alerter"
	"github.com/kevslinger/funda-scraper/scraper"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags:  append(scraper.Flags, alerter.Flags...),
		Action: scrapeFunda,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func scrapeFunda(ctx *cli.Context) error {
	// TODO: Should New take in the actual values instead of a Config object?
	fundaScraper := scraper.New(*scraper.Defaults, &http.Client{})
	resp, err := fundaScraper.Request(http.MethodGet, ctx.String("listing-path"), nil)
	if err != nil {
		log.Fatalf("Error %e", err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error with reading Body: %e", err)
	}
	// TODO: Parse values of the house and store in a DB

	alerts, err := alerter.New(alerter.Config{DiscordAuthenticationToken: ctx.String("discord-auth-token"), DiscordChannelID: ctx.String("discord-channel-id")})
	if err != nil {
		log.Fatalf("Error with creating Alerter: %e", err)
	}
	alerts.Alert("Found a new house at https://funda.nl/" + ctx.String("listing-path"))
	return nil
}
