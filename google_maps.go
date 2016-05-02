package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type GoogleMapsClient struct {
	BaseURL string
}

func NewGoogleMapsClient(c Config) GoogleMapsClient {
	client := GoogleMapsClient{
		BaseURL: "https://maps.googleapis.com/maps/api/distancematrix/json",
	}
	if c.APIBaseURL != "" {
		client.BaseURL = c.APIBaseURL
	}
	return client
}

func (c GoogleMapsClient) GetDuration(start Coordinates, end Coordinates, mode string) (string, error) {
	query := fmt.Sprintf(
		"?origins=%v,%v&&destinations=%v,%v&mode=%v",
		start.Latitude, start.Longitude, end.Latitude,
		end.Longitude, mode,
	)

	resp, err := http.Get(c.BaseURL + query)
	if err != nil {
		return "", errors.New(ErrGettingTimes)
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(ErrGettingTimes)
	}

	var r GoogleMapsResponse

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return "", errors.New(ErrGettingTimes)
	}

	result := r.Rows[0].Elements[0]
	if result.Status != "OK" {
		return "", errors.New(ErrNoDataAvailable)
	}

	durationText := result.Duration.Text

	return durationText, nil
}

type GoogleMapsResponse struct {
	Rows []row `json:"rows"`
}

type row struct {
	Elements []*element `json:"elements"`
}

type element struct {
	Duration duration `json:"duration"`
	Status   string   `json:"status"`
}

type duration struct {
	Text string `json:"text"`
}
