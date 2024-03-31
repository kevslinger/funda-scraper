# Funda Scraper

- [Funda Scraper](#funda-scraper)
  - [Design](#design)
  - [Configuration Parameters](#configuration-parameters)
  - [Funda Scraper](#funda-scraper-1)
  - [Database](#database)
  - [Alerter](#alerter)


## Design

The goal of this project is to scrape funda and send alerts when new listings matching a set of configurable desiderata are found.
The code will contain the following components:

- Configuration parameters 
- Funda scraper
- Database
- Alerter

## Configuration Parameters

The user may specify the criteria they are looking for as if they were on the Funda website.
Examples include whether they want to rent or to buy, which area(s) they are looking for, the min and/or max price they are willing to pay, the min and/or max size they are looking for, number of (bed)rooms, etc.

## Funda Scraper

This component is responsible for communicating with Funda via HTTP requests.
It will gather the requirements from the configuration parameters, pass those in as inputs to the HTTP request, and parse the response. 

## Database

The database will store the results in a SQL DB. 
This will be particularly useful for determining if an entry has already been part of an alert message.

DB table will have the following columns:

- Time seen
- Link (Is this useful?)
- house id (does it exist??)
- Address
- Price
- Description
- Listed Since
- Zip code
- Bouwjaar
- Size (square meters)
- Living Area
- Kind of house (e.g. woonbouw, appartement)
- Building type 
- Number of rooms
- Number of bedrooms
- Layout (what)
- Energy Label
- Insulation (what)
- Heating
- Ownership
- Exteriors
- Parking
- Neighbourhood name
- Date listed
- Date sold
- Term
- Price sold
- Last ask price
- Last ask price m2  (spilt "\r")[0]


## Alerter

The alerter will be responsible for sending messages to various platforms, such as Slack, Discord, Telegram, etc., based on the user's preferences.
