package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
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
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.183 Safari/537.36")
	})
	return &Scraper{
		client:    client,
		collector: collector,
		config:    config,
	}
}

func (s Scraper) GetListingUrls(requestType string, body io.Reader) ([]string, error) {
	// TODO: More options, less hardcoded
	urlPath, err := s.configureUrl()
	if err != nil {
		return nil, err
	}
	return s.getUrlsFromRequest(urlPath, requestType, body)
}

func (s Scraper) configureUrl() (string, error) {
	path, err := url.JoinPath(s.config.baseUrl, s.config.koopOrHuur)
	if err != nil {
		return "", fmt.Errorf("error joining paths %s and %s: %e", s.config.baseUrl, s.config.koopOrHuur, err)
	}
	path += "?"
	if s.config.area != nil {
		path += fmt.Sprintf("selected_area=[\"%s\"]", strings.Join(s.config.area, "\",\""))
	}

	path += s.getPathComponentForIntRange(s.config.minPrice, s.config.maxPrice, "price")
	path += s.getPathComponentForIntRange(s.config.minLivingArea, s.config.maxLivingArea, "floor_area")
	path += s.getPathComponentForIntRange(s.config.minPlotArea, s.config.maxPlotArea, "plot_area")
	path += s.getPathComponentForIntRange(s.config.minRooms, s.config.maxRooms, "rooms")
	path += s.getPathComponentForIntRange(s.config.minBedrooms, s.config.maxBedrooms, "bedrooms")
	// TODO: Add more arguments

	log.Print("Full path is ", "path=", path)
	return path, nil
}

func (s Scraper) getPathComponentForIntRange(lower, upper int, componentName string) string {
	if lower == 0 && upper == 0 {
		return ""
	}
	path := fmt.Sprintf("&%s=\"", componentName)
	if lower != 0 {
		path += fmt.Sprint(lower)
	}
	path += "-"
	if upper != 0 {
		path += fmt.Sprint(upper)
	}
	return path + "\""
}
func (s Scraper) getUrlsFromRequest(fullPath, requestType string, body io.Reader) ([]string, error) {
	var urls []string
	s.collector.OnHTML("script", func(e *colly.HTMLElement) {
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
				log.Print("Url not found")
			}
			log.Print(url)
			urls = append(urls, url.(string))
			// TODO: Create a queue for something to consume the URLs?
		}
	})
	err := s.collector.Request(requestType, fullPath, body, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP %s request with path %s: %e", requestType, fullPath, err)
	}
	return urls, nil
}

// TODO:
func GetFundaListingFromUrl(url string) (FundaListing, error) {
	return FundaListing{}, nil
}
