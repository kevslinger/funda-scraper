package scraper

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Scraper struct {
	client *http.Client
	config Config
}

type FundaListing struct{}

func New(config Config, client *http.Client) *Scraper {
	return &Scraper{
		client: client,
		config: config,
	}
}

func (s Scraper) Request(requestType, path string, body io.Reader) (*http.Response, error) {
	fullPath, err := url.JoinPath(s.config.base_url, path)
	if err != nil {
		return &http.Response{}, fmt.Errorf("error joining paths %s and %s: %e", s.config.base_url, path, err)
	}
	req, err := http.NewRequest(requestType, fullPath, body)
	if err != nil {
		return &http.Response{}, fmt.Errorf("error creating HTTP %s request with path %s: %e", requestType, path, err)
	}
	for k, v := range s.config.headers {
		req.Header.Set(k, v)
	}
	return s.client.Do(req)
}

func (s Scraper) ParseResponse(response *http.Response) FundaListing {
	return FundaListing{}
}
