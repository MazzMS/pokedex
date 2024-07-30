package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "location/"
	fullURL := baseURL + endpoint

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

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreas := LocationAreasResp{}
	err = json.Unmarshal(body, &locationAreas)

	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}
