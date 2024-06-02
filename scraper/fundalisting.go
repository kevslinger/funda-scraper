package scraper

type FundaListing struct {
	URL                  string
	HouseID              string
	Address              string
	Price                int
	Description          string
	ListedSince          string // TODO: Timestamp?
	Postcode             string
	BuildYear            int
	TotalSize            int
	LivingSize           int
	HouseType            string
	BuildingType         string
	NumRooms             int
	NumBedrooms          int
	Layout               string
	EnergyLabel          string
	Insulation           string
	Heating              string
	OwnershipType        string
	Exteriors            string
	Parking              string
	Neighboourhood       string
	DateListed           string // TODO: Timestamp?
	DateSold             string // TODO: Timestamp?
	Term                 int
	PriceSold            int
	LastAskPrice         int
	LastAskPricePerMeter int
}
