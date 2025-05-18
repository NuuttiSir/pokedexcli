package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	val, ok := c.cache.Get(url)
	if ok {
		locationsRes := Location{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return Location{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
