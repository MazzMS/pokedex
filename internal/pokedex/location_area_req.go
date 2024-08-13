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
	section := "/location-area"
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
	// I should check the tradeoff difference between the memory usage of
	// LocationsAreaResp vs []byte and the speed improvement of not having to
	// Unmarshal the json each time.
	locationAreas := LocationAreasResp{}
	err := json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}


func (c *Client) LocationArea(area string) (LocationArea, error) {
	// in order to be efficient with the cache, either we use the id or the name
	// of the location, if not, we could have the same cache twice using id and name

	// to do so, if the user provides a number, we look through the locations
	// and get the name of the location, then use it to do the fetch

	var areaName string
	// check if it is a number
	intId, err := strconv.Atoi(area)
	// if is a number, therefore not an error, get name
	if err == nil {
		// get offset
		offset := intId - (intId % 20)
		if offset < 0 {
			offset = 0
		}

		// get location area/s
		locationAreas, err := c.ListLocationAreas(offset)
		if err != nil {
			return LocationArea{}, err
		}
		// get name through results
		areaName = locationAreas.Results[intId%20 - 1].Name
	} else {
		// if it was already a string, then it is a name 
		areaName = area
	}

	path := "location-area/" + areaName
	fullURL := baseURL + path

	// check if chached
	body, ok := c.cache.Get(fullURL)
	// if not cached
	if !ok {
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return LocationArea{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationArea{}, err
		}
		if resp.StatusCode > 399 {
			return LocationArea{}, fmt.Errorf("Response failed with status code: %d", resp.StatusCode)
		}

		body, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return LocationArea{}, err
		}

		c.cache.Add(fullURL, body)
	}
	// I should check the tradeoff difference between the memory usage of
	// LocationsAreaResp vs []byte and the speed improvement of not having to
	// Unmarshal the json each time.
	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
