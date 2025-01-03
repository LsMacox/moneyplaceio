package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Client struct {
	BaseURL string
	client  *http.Client
	Token   string
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		BaseURL: baseURL,
		client:  &http.Client{},
		Token:   token,
	}
}

func (c *Client) PerformRequest(method, path string, params interface{}) (*http.Response, error) {
	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s%s?%s", c.BaseURL, path, queryParams.Encode())

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Token "+c.Token)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	return resp, nil
}

func (c *Client) DecodeResponse(resp *http.Response, target any) error {
	return json.NewDecoder(resp.Body).Decode(target)
}
