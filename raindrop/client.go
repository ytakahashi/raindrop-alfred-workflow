package raindrop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

const endpoint string = "https://api.raindrop.io"

// Client is a raindrop client
type Client struct {
	baseURL     *url.URL
	httpClient  *http.Client
	accessToken string
}

// Collection represents get collections api response item
type Collection struct {
	ID         uint32 `json:"_id"`
	Color      string `json:"color"`
	Count      uint32 `json:"count"`
	Created    string `json:"created"`
	LastUpdate string `json:"lastUpdate"`
	Expanded   bool   `json:"expanded"`
	Public     bool   `json:"public"`
	Title      string `json:"title"`
	View       string `json:"view"`
}

// Collections represents get collections api response
type Collections struct {
	Result bool         `json:"result"`
	Items  []Collection `json:"items"`
}

// Raindrop represents get raindrops api response item
type Raindrop struct {
	Tags    []string `json:"tags"`
	Cover   string   `json:"cover"`
	Type    string   `json:"type"`
	HTML    string   `json:"html"`
	Excerpt string   `json:"excerpt"`
	Title   string   `json:"title"`
	Link    string   `json:"link"`
}

// Raindrops represents get raindrops api response
type Raindrops struct {
	Result bool       `json:"result"`
	Items  []Raindrop `json:"items"`
}

// NewClient creates Raindrop Client
func NewClient(accessToken string) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	client := Client{
		baseURL:     u,
		httpClient:  &http.Client{},
		accessToken: accessToken,
	}

	return &client, nil
}

// GetCollections call Get root collections API (ref. https://developer.raindrop.io/v1/collections/methods#get-root-collections)
func (c *Client) GetCollections() (*Collections, error) {
	path := "/rest/v1/collections"
	req, err := c.newRequest("GET", path)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	r := new(Collections)
	if err := parseResponse(response, 200, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetRaindrops call Get raindrops API (refs. https://developer.raindrop.io/v1/raindrops/multiple#get-raindrops)
func (c *Client) GetRaindrops(collectionID string) (*Raindrops, error) {
	path := fmt.Sprintf("/rest/v1/raindrops/%s", collectionID)
	req, err := c.newRequest("GET", path)
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	r := new(Raindrops)
	if err := parseResponse(response, 200, &r); err != nil {
		return nil, err
	}

	return r, nil
}

func (c *Client) newRequest(method string, apiPath string) (*http.Request, error) {
	u := *c.baseURL
	u.Path = path.Join(c.baseURL.Path, apiPath)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", c.accessToken)
	req.Header.Add("Authorization", bearerToken)

	return req, nil
}

func parseResponse(response *http.Response, expectedStatus int, clazz interface{}) error {
	if response.StatusCode != expectedStatus {
		return fmt.Errorf("Unexpected Status Code: %d", response.StatusCode)
	}

	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(clazz)
}
