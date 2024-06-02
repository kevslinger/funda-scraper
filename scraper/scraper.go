package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/kevslinger/funda-scraper/config"
)

var (
	urlQueue      []string
	recentListing *FundaListing
)

type Scraper struct {
	client    *http.Client
	collector *colly.Collector
	config    config.ScraperConfig
}

func New(config config.ScraperConfig, client *http.Client) *Scraper {
	collector := colly.NewCollector()
	collector.AllowURLRevisit = true
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
			slog.Error("HTML Element type assertion error", "HTML element", elem)
			return
		}
		for _, e := range elemSlice {
			eMap, ok := e.(map[string]any)
			if !ok {
				slog.Error("HTML Element type assertion error", "HTML element", e)
				return
			}
			url, ok := eMap["url"]
			if !ok {
				return
			}
			urlQueue = append(urlQueue, url.(string))
		}
	})
	scraper.collector.OnHTML("body", func(e *colly.HTMLElement) {
		recentListing.Address = e.ChildText(".object-header__container .text-2xl")
		recentListing.Price = convertStrToInt(e.ChildText("div.gap-2 > span:nth-child(1)"))
		recentListing.Description = e.ChildText(".listing-description-text")
		recentListing.Postcode = e.ChildText(".object-header__container .text-neutral-40")
		recentListing.BuildYear = convertStrToInt(e.ChildText("section.mt-6 > div:nth-child(3) > dl:nth-child(2) > div:nth-child(3) > dd:nth-child(2) > span:nth-child(1)"))
		recentListing.TotalSize = convertStrToInt(e.ChildText("section.mt-6 > div:nth-child(4) > dl:nth-child(2) > div:nth-child(2) > dd:nth-child(2) > span:nth-child(1)"))
		recentListing.LivingSize = convertStrToInt(e.ChildText("section.mt-6 > div:nth-child(4) > dl:nth-child(2) > div:nth-child(1) > dd:nth-child(2) > dl:nth-child(1) > div:nth-child(1) > dd:nth-child(2)"))
		recentListing.HouseType = e.ChildText("section.mt-6 > div:nth-child(3) > dl:nth-child(2) > div:nth-child(1) > dd:nth-child(2) > span:nth-child(1)")
		recentListing.BuildingType = e.ChildText("section.mt-6 > div:nth-child(3) > dl:nth-child(2) > div:nth-child(2) > dd:nth-child(2) > span:nth-child(1)")
		recentListing.NumRooms = convertStrToInt(e.ChildText("section.mt-6 > div:nth-child(5) > dl:nth-child(2) > div:nth-child(1) > dd:nth-child(2) > span:nth-child(1)"))
		recentListing.NumBedrooms = convertStrToInt(e.ChildText("ul.mt-2 > li:nth-child(2) > span:nth-child(2)"))
		// TODO: Add more fields
	})

	return scraper
}

// TODO: Get multiple pages of listings
func (s Scraper) GetListingUrls(requestType string, body io.Reader) ([]string, error) {
	urlPath, err := s.configureUrl()
	if err != nil {
		return nil, err
	}
	urls, err := s.getUrlsFromRequest(urlPath, requestType, body)
	return urls, err
}

func (s Scraper) configureUrl() (string, error) {
	path, err := url.JoinPath(s.config.BaseUrl, s.config.KoopOrHuur)
	if err != nil {
		return "", fmt.Errorf("error joining paths %s and %s: %e", s.config.BaseUrl, s.config.KoopOrHuur, err)
	}
	path += "?"
	path += s.getPathComponentForStringSlice(s.config.Area, "selected_area")
	path += s.getPathComponentForIntRange(s.config.MinPrice, s.config.MaxPrice, "&price")
	path += s.getPathComponentForIntRange(s.config.MinLivingArea, s.config.MaxLivingArea, "&floor_area")
	path += s.getPathComponentForIntRange(s.config.MinPlotArea, s.config.MaxPlotArea, "&plot_area")
	path += s.getPathComponentForIntRange(s.config.MinRooms, s.config.MaxRooms, "&rooms")
	path += s.getPathComponentForIntRange(s.config.MinBedrooms, s.config.MaxBedrooms, "&bedrooms")
	path += s.getPathComponentForStringSlice(s.config.EnergyLabel, "&energy_label")
	path += s.getPathComponentForStringSlice(s.config.OutdoorAmenities, "&exterior_space_type")
	path += s.getPathComponentForStringSlice(s.config.GardenDirection, "&exterior_space_garden_orientation")
	path += s.getPathComponentForIntRange(s.config.GardenMinSpace, s.config.GardenMaxSpace, "&exterior_space_garden_size")
	path += s.getPathComponentForStringSlice(s.config.BuildingType, "&construction_type")
	path += s.getPathComponentForStringSlice(s.config.Zoning, "zoning")
	path += s.getPathComponentForStringSlice(s.config.ConstructionPeriod, "&construction_period")
	path += s.getPathComponentForStringSlice(s.config.Surroundings, "&surrounding")
	path += s.getPathComponentForStringSlice(s.config.Garage, "&garage_type")
	path += s.getPathComponentForIntRange(s.config.MinGarageCapacity, s.config.MaxGarageCapacity, "&garage_capacity")
	path += s.getPathComponentForStringSlice(s.config.Characteristics, "&amenities")
	path += s.getPathComponentForStringSlice(s.config.Display, "&type")
	path += s.getPathComponentForStringSlice(s.config.OpenHouse, "&open_house")

	// TODO: Choose other sort methods?
	path += "&sort=\"date_down\""
	slog.Info("Full search path is", "path", path)
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

// TODO: Add more fields
func (s Scraper) GetFundaListingFromUrl(url string) (FundaListing, error) {
	err := s.collector.Visit(url)
	if err != nil {
		return FundaListing{}, err
	}
	recentListing.URL = url
	return *recentListing, nil
}

func convertStrToInt(in string) int {
	reg, _ := regexp.Compile("[^0-9]+")
	numStr := reg.ReplaceAllString(in, "")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return -1
	}
	return num
}
