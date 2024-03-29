package main

func main() {

}

/*
Main is going to be responsible for loading in the config,
starting up the components, and running the application.

Component hierarchy:
Config depends on: nothing
Alerter depends on: config
Database depends on: config
Scraper depends on: config, database, alerter
main depends on: config, database, alerter, scraper

The scraper will maintain the flow of the application
*/
