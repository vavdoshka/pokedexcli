package pokeapi

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

	locationResponse, err := genericRequestWithCache[LocationAreaResponse](c, url)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationResponse, nil
}
