package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	dataURL = "data"
	dIdxURL = "daily-index"
)

// New creates new EDGAR http client
func New() *Client {
	return &Client{
		baseURL: "https://www.sec.gov/Archives/edgar/",
	}
}

// Client represents EDGAR client
type Client struct {
	baseURL string
}

// RequestDataContent accepts the params bla bla
func (c *Client) RequestDataContent(params []string) ([]byte, error) {
	filingURL := c.baseURL + dataURL
	for _, param := range params {
		filingURL += "/" + param
	}

	resp, err := http.Get(filingURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// RequestDataItems accepts the params bla bla
func (c *Client) RequestDataItems(params []string) ([]Item, error) {
	filingURL := c.baseURL + dataURL
	for _, param := range params {
		filingURL += "/" + param
	}
	filingURL += ".json"

	resp, err := c.request(filingURL)
	if err != nil {
		return nil, err
	}

	return resp.Directory.Items, nil
}

// RequestDailyIndexItems blabla
func (c *Client) RequestDailyIndexItems(params []string) ([]Item, error) {
	filingURL := c.baseURL + dIdxURL
	for _, param := range params {
		filingURL += "/" + param
	}
	filingURL += ".json"

	resp, err := c.request(filingURL)
	if err != nil {
		return nil, err
	}

	return resp.Directory.Items, nil
}

func (c *Client) request(url string) (*Response, error) {
	hresp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer hresp.Body.Close()
	body, err := ioutil.ReadAll(hresp.Body)
	if err != nil {
		return nil, err
	}

	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
