package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

type locationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetPokemons(url string) (locationAreaResponse, error) {
	var result = locationAreaResponse{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	bytesRes, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	json.Unmarshal(bytesRes, &result)

	return result, nil
}
