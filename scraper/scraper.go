package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gocolly/colly"
)

var (
	urlQueue      []string
	recentListing *FundaListing
)

type Scraper struct {
	client    *http.Client
	collector *colly.Collector
	config    Config
}

type FundaListing struct {
	URL     string
	Address string
}

func New(config Config, client *http.Client) *Scraper {
	collector := colly.NewCollector()
	scraper := &Scraper{
		client:    client,
		collector: collector,
		config:    config,
	}
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
		recentListing = &FundaListing{}
		urlQueue = make([]string, 0)
	})

	scraper.collector.OnHTML("script", func(e *colly.HTMLElement) {
		// Marshal the json into a generic map
		result := make(map[string]any)
		json.Unmarshal([]byte(e.Text), &result)
		// Filter out all other scripts that don't have itemListElement
		// TODO: We know the script has type "application/ld+json"
		// TODO: Can we use that fact to make an easier filter?
		elem, ok := result["itemListElement"]
		if !ok {
			return
		}
		// TODO: Don't use Fatalf, just print a warning and return
		elemSlice, ok := elem.([]any)
		if !ok {
			// TODO: Better print format?
			log.Fatalf("error with type assertion. HTML element=%#v", elem)
		}
		for _, e := range elemSlice {
			eMap, ok := e.(map[string]any)
			if !ok {
				log.Fatalf("error with type assertion. HTML element=%#v", e)
			}
			url, ok := eMap["url"]
			if !ok {
				return
			}
			log.Print(url)
			urlQueue = append(urlQueue, url.(string))
		}
	})
	scraper.collector.OnHTML("body", func(e *colly.HTMLElement) {
		recentListing.Address = e.ChildText(".object-header__title")
		// TODO: Add more fields
	})

	return scraper
}

func (s Scraper) GetListingUrls(requestType string, body io.Reader) ([]string, error) {
	urlPath, err := s.configureUrl()
	if err != nil {
		return nil, err
	}
	urls, err := s.getUrlsFromRequest(urlPath, requestType, body)
	return urls, err
}

func (s Scraper) configureUrl() (string, error) {
	path, err := url.JoinPath(s.config.baseUrl, s.config.koopOrHuur)
	if err != nil {
		return "", fmt.Errorf("error joining paths %s and %s: %e", s.config.baseUrl, s.config.koopOrHuur, err)
	}
	path += "?"
	path += s.getPathComponentForStringSlice(s.config.area, "selected_area")
	path += s.getPathComponentForIntRange(s.config.minPrice, s.config.maxPrice, "&price")
	path += s.getPathComponentForIntRange(s.config.minLivingArea, s.config.maxLivingArea, "&floor_area")
	path += s.getPathComponentForIntRange(s.config.minPlotArea, s.config.maxPlotArea, "&plot_area")
	path += s.getPathComponentForIntRange(s.config.minRooms, s.config.maxRooms, "&rooms")
	path += s.getPathComponentForIntRange(s.config.minBedrooms, s.config.maxBedrooms, "&bedrooms")
	path += s.getPathComponentForStringSlice(s.config.energyLabel, "&energy_label")
	path += s.getPathComponentForStringSlice(s.config.outdoorAmenities, "&exterior_space_type")
	path += s.getPathComponentForStringSlice(s.config.gardenDirection, "&exterior_space_garden_orientation")
	path += s.getPathComponentForIntRange(s.config.gardenMinSpace, s.config.gardenMaxSpace, "&exterior_space_garden_size")
	path += s.getPathComponentForStringSlice(s.config.buildingType, "&construction_type")
	path += s.getPathComponentForStringSlice(s.config.zoning, "zoning")
	path += s.getPathComponentForStringSlice(s.config.constructionPeriod, "&construction_period")
	path += s.getPathComponentForStringSlice(s.config.surroundings, "&surrounding")
	path += s.getPathComponentForStringSlice(s.config.garage, "&garage_type")
	path += s.getPathComponentForIntRange(s.config.minGarageCapacity, s.config.maxGarageCapacity, "&garage_capacity")
	path += s.getPathComponentForStringSlice(s.config.characteristics, "&amenities")
	path += s.getPathComponentForStringSlice(s.config.display, "&type")
	path += s.getPathComponentForStringSlice(s.config.openHouse, "&open_house")

	// TODO: Choose other sort methods?
	path += "&sort=\"date_down\""
	log.Print("Full search path is ", "path=", path)
	return path, nil
}

func (s Scraper) getPathComponentForIntRange(lower, upper int, componentName string) string {
	if lower == 0 && upper == 0 {
		return ""
	}
	path := fmt.Sprintf("%s=\"", componentName)
	if lower != 0 {
		path += fmt.Sprint(lower)
	}
	path += "-"
	if upper != 0 {
		path += fmt.Sprint(upper)
	}
	return path + "\""
}

func (s Scraper) getPathComponentForStringSlice(values []string, componentName string) string {
	if len(values) <= 0 {
		return ""
	}
	path := fmt.Sprintf("%s=[", componentName)
	for i, value := range values {
		// Replaces + in energy label with %2B
		path += "\"" + url.QueryEscape(value) + "\""
		if i < len(values)-1 {
			path += ","
		}
	}
	return path + "]"
}

func (s Scraper) getUrlsFromRequest(fullPath, requestType string, body io.Reader) ([]string, error) {
	err := s.collector.Request(requestType, fullPath, body, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP %s request with path %s: %e", requestType, fullPath, err)
	}
	return urlQueue, nil
}

// TODO:
func (s Scraper) GetFundaListingFromUrl(url string) (FundaListing, error) {
	err := s.collector.Visit(url)
	if err != nil {
		return FundaListing{}, err
	}
	recentListing.URL = url
	return *recentListing, nil
}
