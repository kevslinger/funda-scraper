package scraper

import (
	"github.com/urfave/cli/v2"
)

type Config struct {
	headers            map[string][]string
	baseUrl            string
	koopOrHuur         string
	area               []string
	minPrice           int
	maxPrice           int
	minLivingArea      int
	maxLivingArea      int
	minPlotArea        int
	maxPlotArea        int
	minRooms           int
	maxRooms           int
	minBedrooms        int
	maxBedrooms        int
	energyLabel        []string
	outdoorAmenities   []string
	gardenDirection    []string
	gardenMinSpace     int
	gardenMaxSpace     int
	buildingType       []string
	zoning             []string
	constructionPeriod []string
}

const (
	defaultBaseUrl         = "https://www.funda.nl/zoeken"
	koopOrHuurFlag         = "koop-or-huur"
	baseUrlFlag            = "base-url"
	areaFlag               = "area"
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
	zoningFlag             = "zoning"
	constructionPeriodFlag = "construction-period"
)

var Defaults *Config = &Config{
	headers:    map[string][]string{"user-agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36"}},
	baseUrl:    defaultBaseUrl,
	koopOrHuur: "koop",
	area:       []string{"nl"},
}

var Flags []cli.Flag = []cli.Flag{
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
	// TODO: Add Ligging, Garage, Garagecapaciteit, Eigenschappen, Weergave, Open Huis
}

func LoadConfig(ctx *cli.Context) *Config {
	config := Defaults
	if ctx.IsSet(baseUrlFlag) {
		config.baseUrl = ctx.String(baseUrlFlag)
	}
	if ctx.IsSet(koopOrHuurFlag) {
		config.koopOrHuur = ctx.String(koopOrHuurFlag)
	}
	if ctx.IsSet(areaFlag) {
		config.area = ctx.StringSlice(areaFlag)
	}
	if ctx.IsSet(minPriceFlag) {
		config.minPrice = ctx.Int(minPriceFlag)
	}
	if ctx.IsSet(maxPriceFlag) {
		config.maxPrice = ctx.Int(maxPriceFlag)
	}
	if ctx.IsSet(minLivingAreaFlag) {
		config.minLivingArea = ctx.Int(minLivingAreaFlag)
	}
	if ctx.IsSet(maxLivingAreaFlag) {
		config.maxLivingArea = ctx.Int(maxLivingAreaFlag)
	}
	if ctx.IsSet(minPlotAreaFlag) {
		config.minPlotArea = ctx.Int(minPlotAreaFlag)
	}
	if ctx.IsSet(maxPlotAreaFlag) {
		config.maxPlotArea = ctx.Int(maxPlotAreaFlag)
	}
	if ctx.IsSet(minRoomsFlag) {
		config.minRooms = ctx.Int(minRoomsFlag)
	}
	if ctx.IsSet(maxRoomsFlag) {
		config.maxRooms = ctx.Int(maxRoomsFlag)
	}
	if ctx.IsSet(minBedroomsFlag) {
		config.minBedrooms = ctx.Int(minBedroomsFlag)
	}
	if ctx.IsSet(maxBedroomsFlag) {
		config.maxBedrooms = ctx.Int(maxBedroomsFlag)
	}
	if ctx.IsSet(energyLabelFlag) {
		config.energyLabel = ctx.StringSlice(energyLabelFlag)
	}
	if ctx.IsSet(outdooramenitiesFlag) {
		config.outdoorAmenities = ctx.StringSlice(outdooramenitiesFlag)
	}
	if ctx.IsSet(gardenDirectionFlag) {
		config.gardenDirection = ctx.StringSlice(gardenDirectionFlag)
	}
	if ctx.IsSet(gardenMinSpaceFlag) {
		config.gardenMinSpace = ctx.Int(gardenMinSpaceFlag)
	}
	if ctx.IsSet(gardenMaxSpaceFlag) {
		config.gardenMaxSpace = ctx.Int(gardenMaxSpaceFlag)
	}
	if ctx.IsSet(buildingTypeFlag) {
		config.buildingType = ctx.StringSlice(buildingTypeFlag)
	}
	if ctx.IsSet(zoningFlag) {
		config.zoning = ctx.StringSlice(zoningFlag)
	}
	if ctx.IsSet(constructionPeriodFlag) {
		config.constructionPeriod = ctx.StringSlice(constructionPeriodFlag)
	}
	return config
}
