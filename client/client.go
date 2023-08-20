package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseAddress = "http://statsapi.web.nhl.com/api/v1"

// Client for API calls
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: baseAddress,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c Client) FormatURL(endpoint string) string {
	return fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
}

// Make a get request and return any errors if they exist
// Return the body if successful
func (c *Client) GetRequest(endpoint string) ([]byte, error) {
	URL := c.FormatURL(endpoint)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %q", URL)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status %q", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	return data, nil
}
