package honeycombiosdk

import (
	"net/http"
)

const (
	apiURL    = "https://api.honeycomb.io"
	userAgent = "terraform-provider-honeycombio"
)

type Client struct {
	apiKey     string
	dataset    string
	httpClient *http.Client
}

func NewClient(apiKey, dataset string) *Client {
	return &Client{
		apiKey:     apiKey,
		dataset:    dataset,
		httpClient: &http.Client{},
	}
}

func (c *Client) populateHeaders(req *http.Request) {
	req.Header.Add("X-Honeycomb-Team", c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", userAgent)
}