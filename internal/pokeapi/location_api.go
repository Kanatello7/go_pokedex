package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) getAndCache(url string, target any) error {
	if cached, ok := c.cache.Get(url); ok {
		return json.Unmarshal(cached, target)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(dat, target); err != nil {
		return err
	}
	c.cache.Add(url, dat)
	return nil
}

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	var resp RespShallowLocations
	err := c.getAndCache(url, &resp)
	return resp, err
}

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	var resp Location
	err := c.getAndCache(url, &resp)
	return resp, err

}
