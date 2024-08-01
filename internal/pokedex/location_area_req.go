package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (c *Client) ListLocationAreas(offset int) (LocationAreasResp, error) {
	queryParameter := "/?offset=" + strconv.Itoa(offset) + "&limit=20"
	section := "/location"
	fullURL := baseURL + section + queryParameter

	// check if chached
	body, ok := c.cache.Get(fullURL)
	// if not cached
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationAreasResp{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreasResp{}, err
		}
		if resp.StatusCode > 399 {
			return LocationAreasResp{}, fmt.Errorf("Response failed with status code: %d\n", resp.StatusCode)
		}

		body, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return LocationAreasResp{}, err
		}

		c.cache.Add(fullURL, body)
	}
	locationAreas := LocationAreasResp{}
	err := json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}
