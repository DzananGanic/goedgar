package edgar

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/goedgar/edgar/filings"
)

// New creates new EDGAR client
func New() *Client {
	return &Client{
		baseURL: "https://www.sec.gov/Archives/edgar/data",
	}
}

// Client represents EDGAR client
type Client struct {
	baseURL string
}

// FilingsForCIK returns the filings for the company
// with provided CIK number
func (c *Client) FilingsForCIK(cik string) ([]filings.Item, error) {
	filingsURL := c.baseURL + "/" + cik + "/" + "index.json"

	resp, err := http.Get(filingsURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fResp filings.Resp
	err = json.Unmarshal(body, &fResp)
	if err != nil {
		return nil, err
	}

	return fResp.Dir.Items, nil
}

// FilingDocs accepts company cik number,
// filing name and returns filing documents
func (c *Client) FilingDocs(cik, name string) ([]filings.Item, error) {
	filingURL := c.baseURL + "/" + cik + "/" + name + "/" + "index.json"

	resp, err := http.Get(filingURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fResp filings.Resp
	err = json.Unmarshal(body, &fResp)
	if err != nil {
		return nil, err
	}

	return fResp.Dir.Items, nil
}
