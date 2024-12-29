package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Client) ListLocations(pageUrl *string) (LocationAreaResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area"

	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("can not read the response %v, err %w", resp.Body, err)
	}

	locationAreaResponse := LocationAreaResponse{}

	if err := json.Unmarshal(dat, &locationAreaResponse); err != nil {
		return locationAreaResponse, fmt.Errorf("can not unmarshal response %v, err %w", dat, err)
	}

	return locationAreaResponse, nil
}
