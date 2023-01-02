package api

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	DomainUS = "com"
	DomainUK = "co.uk"
	DomainDE = "de"
	DomainFR = "fr"
	DomainJP = "co.jp"
	DomainCA = "ca"
	DomainIT = "it"
	DomainES = "es"
	DomainMX = "com.mx"
	DomainIN = "in"
)

type Client struct {
	ApiKey string
	BaseUrl string
	Client *http.Client
}

func New(apikey string) *Client {
	return &Client{
		ApiKey: apikey,
		BaseUrl: "https://api.keepa.com",
		Client: http.DefaultClient,
	}
}

func (c *Client) post(path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(req)
}

func (c *Client) GetGraph(productID string,domain string) ([]byte, error) {
	paylaod := url.Values{
		"key": {c.ApiKey},
		"domain": {domain},
		"asin": {productID},
	}
	url := c.BaseUrl + "/graphimage?" + paylaod.Encode()

	resp, err := c.post(url, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	html,err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return html, nil
}
