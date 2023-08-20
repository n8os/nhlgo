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
	// // conditions, err := ParseResponse(data)
	// // if err != nil {
	// // 	return Conditions{}, err
	// // }
	// // var result interface{}
	// err := json.Unmarshal(data, &result)
	// if err != nil {
	// 	return schema{}, fmt.Errorf("invalid API response %s: %w", data, err)
	// }
	return data, nil
}

// func (c *Client) getRequest(endpoint string, params map[string]string, schema interface{}) int {
// 	body, status := c.makeRequestWithoutJson(endpoint, params)
// 	json.Unmarshal(body, schema)
// 	return status
// }

// func (c *Client) makeRequestWithoutJson(endpoint string, params map[string]string) ([]byte, int) {
// 	request, _ := http.NewRequest("GET", c.baseURL+endpoint, nil)
// 	request.Header.Set("Content-Type", "application/json")
// 	query := request.URL.Query()
// 	for key, value := range params {
// 		query.Add(key, value)
// 	}
// 	request.URL.RawQuery = query.Encode()
// 	//fmt.Println(request.URL)
// 	response, _ := c.httpClient.Do(request)
// 	//check(err)
// 	defer response.Body.Close()
// 	body, _ := ioutil.ReadAll(response.Body)
// 	return body, response.StatusCode
// }
