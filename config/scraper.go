package config

import (
	"github.com/urfave/cli/v2"
)

type ScraperConfig struct {
	Headers            map[string][]string
	BaseUrl            string
	KoopOrHuur         string
	Area               []string
	Postcode           []string
	MinPrice           int
	MaxPrice           int
	MinLivingArea      int
	MaxLivingArea      int
	MinPlotArea        int
	MaxPlotArea        int
	MinRooms           int
	MaxRooms           int
	MinBedrooms        int
	MaxBedrooms        int
	EnergyLabel        []string
	OutdoorAmenities   []string
	GardenDirection    []string
	GardenMinSpace     int
	GardenMaxSpace     int
	BuildingType       []string
	Zoning             []string
	ConstructionPeriod []string
	Surroundings       []string
	Garage             []string
	MinGarageCapacity  int
	MaxGarageCapacity  int
	Characteristics    []string
	Display            []string
	OpenHouse          []string
}

const (
	defaultBaseUrl         = "https://www.funda.nl/zoeken"
	koopOrHuurFlag         = "koop-or-huur"
	baseUrlFlag            = "base-url"
	areaFlag               = "house-area"
	postcodeFlag           = "postcode"
	minPriceFlag           = "min-price"
	maxPriceFlag           = "max-price"
	minLivingAreaFlag      = "min-living-area"
	maxLivingAreaFlag      = "max-living-area"
	minPlotAreaFlag        = "min-plot-area"
	maxPlotAreaFlag        = "max-plot-area"
	minRoomsFlag           = "min-rooms"
	maxRoomsFlag           = "max-rooms"
	minBedroomsFlag        = "min-bedrooms"
	maxBedroomsFlag        = "max-bedrooms"
	energyLabelFlag        = "energy-label"
	outdooramenitiesFlag   = "outdoor-amenities"
	gardenDirectionFlag    = "garden-direction"
	gardenMinSpaceFlag     = "garden-min-space"
	gardenMaxSpaceFlag     = "garden-max-space"
	buildingTypeFlag       = "building-type"
	zoningFlag             = "house-zoning"
	constructionPeriodFlag = "construction-period"
	surroundingsFlag       = "surroundings"
	garageFlag             = "house-garage"
	minGarageCapacityFlag  = "min-garage-capacity"
	maxGarageCapacityFlag  = "max-garage-capacity"
	characteristicsFlag    = "characteristics"
	displayFlag            = "house-display"
	openHouseFlag          = "open-house"
)

var ScraperDefaults *ScraperConfig = &ScraperConfig{
	Headers:    map[string][]string{"user-agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"}},
	BaseUrl:    defaultBaseUrl,
	KoopOrHuur: "koop",
	Area:       []string{"nl"},
}

// TODO: Translate the options to those used in the URL
var ScraperFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  baseUrlFlag,
		Value: defaultBaseUrl,
		Usage: "The base URL to scan",
	},
	&cli.StringFlag{
		Name:  koopOrHuurFlag,
		Value: "koop",
		Usage: "Whether to search for houses to buy or rent (koop or huur)",
	},
	&cli.StringSliceFlag{
		Name:  areaFlag,
		Usage: "Use this to filter houses by area(s)",
	},
	&cli.StringSliceFlag{
		Name:  postcodeFlag,
		Usage: "Use this to alert only when the house is in one of the specified postal codes.",
	},
	&cli.IntFlag{
		Name:  minPriceFlag,
		Usage: "The minimum house price to search for",
	},
	&cli.IntFlag{
		Name:  maxPriceFlag,
		Usage: "The maximum house price to search for",
	},
	&cli.IntFlag{
		Name:  minLivingAreaFlag,
		Usage: "Use this to filter out houses that have less living area than desired (in m^2)",
	},
	&cli.IntFlag{
		Name:  maxLivingAreaFlag,
		Usage: "Use this to filter out houses that have more living area than desired (in m^2)",
	},
	&cli.IntFlag{
		Name:  minPlotAreaFlag,
		Usage: "Use this to filter out houses that have less plot area than desired (in m^2)",
	},
	&cli.IntFlag{
		Name:  maxPlotAreaFlag,
		Usage: "Use this to filter out houses that have more plot area than desired (in m^2)",
	},
	&cli.IntFlag{
		Name:  minRoomsFlag,
		Usage: "Use this to filter out houses that have less (total) rooms than desired",
	},
	&cli.IntFlag{
		Name:  maxRoomsFlag,
		Usage: "Use this to filter out houses that have more (total) rooms than desired",
	},
	&cli.IntFlag{
		Name:  minBedroomsFlag,
		Usage: "Use this to filter out houses that have less bedrooms than desired",
	},
	&cli.IntFlag{
		Name:  maxBedroomsFlag,
		Usage: "Use this to filter out houses that have more bedrooms than desired",
	},
	&cli.StringSliceFlag{
		Name:  energyLabelFlag,
		Usage: "Use this to select which energy label(s) are desired. Options: 'A+++++', 'A++++', 'A+++', 'A++', 'A+', 'A', 'B', 'C', 'D', 'E', 'F', 'G'",
	},
	&cli.StringSliceFlag{
		Name:  outdooramenitiesFlag,
		Usage: "Use this to select which outdoor ammenitie(s) are desired. Options: 'Balkon', 'Dakterras', 'Tuin'",
	},
	&cli.StringSliceFlag{
		Name:  gardenDirectionFlag,
		Usage: "Use this to select which direction(s) the garden can face. Options: 'Zuid', 'West', 'Oost', 'Noord'",
	},
	&cli.IntFlag{
		Name:  gardenMinSpaceFlag,
		Usage: "Use this to filter out houses that have less garden space than desired (in m^2)",
	},
	&cli.IntFlag{
		Name:  gardenMaxSpaceFlag,
		Usage: "Use this to filter out houses that have more garden space than desired (in m^2)",
	},
	&cli.StringSliceFlag{
		Name:  buildingTypeFlag,
		Usage: "Use this to choose between existing buildings or new buildings. Options: 'Bestaande bouw', 'Nieuwbouw'",
	},
	&cli.StringSliceFlag{
		Name:  zoningFlag,
		Usage: "Use this to filter houses based on the purpose. Options: 'Permanente bewoning', 'Recreatiewoning'",
	},
	&cli.StringFlag{
		Name:  constructionPeriodFlag,
		Usage: "Use this to filter houses based on when they were built. Options: 'Voor 1906', '1906-1930', '1931-1944', '1945-1959', '1960-1970', '1971-1980', '1981-1990', '1991-2000', '2001-2010', 2011-2020', 'Onbekend', 'Na 2020'",
	},
	&cli.StringSliceFlag{
		Name:  surroundingsFlag,
		Usage: "Use this to filter houses based on what you would like surrounding your house. Options: 'Aan bosrand', 'Aan water', 'In centrum', 'In bosrijke omgeving', 'In woonwijk', 'Aan drukke weg', 'Aan vaarwater', 'Aan rustige weg', 'Open ligging', 'Buiten bebouwde kom', 'Aan park', 'Landelijk gelegen', 'Beschutte ligging', 'Vrij uitzicht'",
	},
	&cli.StringSliceFlag{
		Name:  garageFlag,
		Usage: "Use this to filter houses based on what type of garage is desired. Options: 'Souterrain', 'Inpandige garage', 'Vrijstaande garage', 'Garage + Carport', 'Aangebouwde garage', 'Garagebox', 'Parkeerkelder'",
	},
	&cli.IntFlag{
		Name:  minGarageCapacityFlag,
		Usage: "Use this to filter houses based on the minimum desired garage capacity",
	},
	&cli.IntFlag{
		Name:  maxGarageCapacityFlag,
		Usage: "Use this to filter houses based on the maximum desired garage capacity",
	},
	&cli.StringSliceFlag{
		Name:  characteristicsFlag,
		Usage: "Use this to filter houses based on certain characteristics. Options: 'Lig-/zitbad', 'CV-ketel', 'Dubbele bewoning', 'Open haard', 'Kluswoning', 'Duurzame energie', 'Zwembad'",
	},
	&cli.StringSliceFlag{
		Name:  displayFlag,
		Usage: "Use this to filter houses based on display. Options: 'Projecten', 'Woningen'",
	},
	&cli.StringSliceFlag{
		Name:  openHouseFlag,
		Usage: "Use this to filter houses based on open house. Options: 'Alle open huizen', 'Komend weekend', 'Vandaag'",
	},
}

func LoadScraperConfig(ctx *cli.Context) *ScraperConfig {
	c := ScraperDefaults
	// Set ENV vars before command-line flags
	readStringFromEnv(baseUrlFlag, &c.BaseUrl)
	readStringFromEnv(koopOrHuurFlag, &c.KoopOrHuur)
	readStringSliceFromEnv(areaFlag, &c.Area)
	readStringSliceFromEnv(postcodeFlag, &c.Postcode)
	readIntFromEnv(minPriceFlag, &c.MinPrice)
	readIntFromEnv(maxPriceFlag, &c.MaxPrice)
	readIntFromEnv(minLivingAreaFlag, &c.MinLivingArea)
	readIntFromEnv(maxLivingAreaFlag, &c.MaxLivingArea)
	readIntFromEnv(minPlotAreaFlag, &c.MinPlotArea)
	readIntFromEnv(maxPlotAreaFlag, &c.MaxPlotArea)
	readIntFromEnv(minRoomsFlag, &c.MinRooms)
	readIntFromEnv(maxRoomsFlag, &c.MaxRooms)
	readIntFromEnv(minBedroomsFlag, &c.MinBedrooms)
	readIntFromEnv(maxBedroomsFlag, &c.MaxBedrooms)
	readStringSliceFromEnv(energyLabelFlag, &c.EnergyLabel)
	readStringSliceFromEnv(outdooramenitiesFlag, &c.OutdoorAmenities)
	readStringSliceFromEnv(gardenDirectionFlag, &c.GardenDirection)
	readIntFromEnv(gardenMinSpaceFlag, &c.GardenMinSpace)
	readIntFromEnv(gardenMaxSpaceFlag, &c.GardenMaxSpace)
	readStringSliceFromEnv(buildingTypeFlag, &c.BuildingType)
	readStringSliceFromEnv(zoningFlag, &c.Zoning)
	readStringSliceFromEnv(constructionPeriodFlag, &c.ConstructionPeriod)
	readStringSliceFromEnv(surroundingsFlag, &c.Surroundings)
	readStringSliceFromEnv(garageFlag, &c.Garage)
	readIntFromEnv(minGarageCapacityFlag, &c.MinGarageCapacity)
	readIntFromEnv(maxGarageCapacityFlag, &c.MaxGarageCapacity)
	readStringSliceFromEnv(characteristicsFlag, &c.Characteristics)
	readStringSliceFromEnv(displayFlag, &c.Display)
	readStringSliceFromEnv(openHouseFlag, &c.OpenHouse)

	if ctx.IsSet(baseUrlFlag) {
		c.BaseUrl = ctx.String(baseUrlFlag)
	}
	if ctx.IsSet(koopOrHuurFlag) {
		c.KoopOrHuur = ctx.String(koopOrHuurFlag)
	}
	if ctx.IsSet(areaFlag) {
		c.Area = ctx.StringSlice(areaFlag)
	}
	if ctx.IsSet(postcodeFlag) {
		c.Postcode = ctx.StringSlice(postcodeFlag)
	}
	if ctx.IsSet(minPriceFlag) {
		c.MinPrice = ctx.Int(minPriceFlag)
	}
	if ctx.IsSet(maxPriceFlag) {
		c.MaxPrice = ctx.Int(maxPriceFlag)
	}
	if ctx.IsSet(minLivingAreaFlag) {
		c.MinLivingArea = ctx.Int(minLivingAreaFlag)
	}
	if ctx.IsSet(maxLivingAreaFlag) {
		c.MaxLivingArea = ctx.Int(maxLivingAreaFlag)
	}
	if ctx.IsSet(minPlotAreaFlag) {
		c.MinPlotArea = ctx.Int(minPlotAreaFlag)
	}
	if ctx.IsSet(maxPlotAreaFlag) {
		c.MaxPlotArea = ctx.Int(maxPlotAreaFlag)
	}
	if ctx.IsSet(minRoomsFlag) {
		c.MinRooms = ctx.Int(minRoomsFlag)
	}
	if ctx.IsSet(maxRoomsFlag) {
		c.MaxRooms = ctx.Int(maxRoomsFlag)
	}
	if ctx.IsSet(minBedroomsFlag) {
		c.MinBedrooms = ctx.Int(minBedroomsFlag)
	}
	if ctx.IsSet(maxBedroomsFlag) {
		c.MaxBedrooms = ctx.Int(maxBedroomsFlag)
	}
	if ctx.IsSet(energyLabelFlag) {
		c.EnergyLabel = ctx.StringSlice(energyLabelFlag)
	}
	if ctx.IsSet(outdooramenitiesFlag) {
		c.OutdoorAmenities = ctx.StringSlice(outdooramenitiesFlag)
	}
	if ctx.IsSet(gardenDirectionFlag) {
		c.GardenDirection = ctx.StringSlice(gardenDirectionFlag)
	}
	if ctx.IsSet(gardenMinSpaceFlag) {
		c.GardenMinSpace = ctx.Int(gardenMinSpaceFlag)
	}
	if ctx.IsSet(gardenMaxSpaceFlag) {
		c.GardenMaxSpace = ctx.Int(gardenMaxSpaceFlag)
	}
	if ctx.IsSet(buildingTypeFlag) {
		c.BuildingType = ctx.StringSlice(buildingTypeFlag)
	}
	if ctx.IsSet(zoningFlag) {
		c.Zoning = ctx.StringSlice(zoningFlag)
	}
	if ctx.IsSet(constructionPeriodFlag) {
		c.ConstructionPeriod = ctx.StringSlice(constructionPeriodFlag)
	}
	if ctx.IsSet(surroundingsFlag) {
		c.Surroundings = ctx.StringSlice(surroundingsFlag)
	}
	if ctx.IsSet(garageFlag) {
		c.Garage = ctx.StringSlice(garageFlag)
	}
	if ctx.IsSet(minGarageCapacityFlag) {
		c.MinGarageCapacity = ctx.Int(minGarageCapacityFlag)
	}
	if ctx.IsSet(maxGarageCapacityFlag) {
		c.MaxGarageCapacity = ctx.Int(maxGarageCapacityFlag)
	}
	if ctx.IsSet(characteristicsFlag) {
		c.Characteristics = ctx.StringSlice(characteristicsFlag)
	}
	if ctx.IsSet(displayFlag) {
		c.Display = ctx.StringSlice(displayFlag)
	}
	if ctx.IsSet(openHouseFlag) {
		c.OpenHouse = ctx.StringSlice(openHouseFlag)
	}
	return c
}
