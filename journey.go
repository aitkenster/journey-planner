package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Journey struct {
	Start Coordinates
	End   Coordinates
	Times journeyTimes
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type journeyTimes struct {
	Car     string `json:"car"`
	Bicycle string `json:"bicycle"`
	Walk    string `json:"walk"`
}

func NewJourney(startVals, endVals string) (*Journey, error) {
	var j Journey
	err := j.Start.toCoordinates(startVals)
	if err != nil {
		return nil, err
	}
	err = j.End.toCoordinates(endVals)
	if err != nil {
		return nil, err
	}
	return &j, nil
}

func (c *Coordinates) toCoordinates(vals string) error {
	var err error
	// TODO what if it's nil?
	splitVals := strings.Split(vals, ",")
	c.Latitude, err = strconv.ParseFloat(splitVals[0], 64)
	if err != nil {
		return err
	}
	c.Longitude, err = strconv.ParseFloat(splitVals[1], 64)
	if err != nil {
		return err
	}
	if c.AreValid() {
		return nil
	}
	return errors.New("invalid coordinates")
}

func (c Coordinates) AreValid() bool {
	return (-90.0 <= c.Latitude && c.Latitude <= 90.0) &&
		(-180.0 <= c.Longitude && c.Longitude <= 180.0)
}

type transportMode int

const (
	carMode = iota
	walkMode
	bicycleMode
)

func (j *Journey) GetTimes() error {
	var err error
	j.Times.Car, err = j.GetTimeFor(carMode)
	if err != nil {
		return err
	}

	j.Times.Walk, err = j.GetTimeFor(walkMode)
	if err != nil {
		return err
	}

	j.Times.Bicycle, err = j.GetTimeFor(bicycleMode)
	if err != nil {
		return err
	}
	return nil
}

func (j *Journey) GetTimeFor(mode transportMode) (string, error) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%v,%v&&destinations=%v,%v&mode=", j.Start.Latitude, j.Start.Longitude, j.End.Latitude, j.End.Longitude)
	var r APIResponse

	switch mode {
	case carMode:
		url = url + "car"
	case walkMode:
		url = url + "walking"
	case bicycleMode:
		url = url + "bicycling"
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	json.NewDecoder(resp.Body).Decode(&r)
	return r.Rows[0].Elements[0].Duration.Text, nil
}

type APIResponse struct {
	Rows []row `json:"rows"`
}

type row struct {
	Elements []*element `json:"elements"`
}

type element struct {
	Duration duration `json:"duration"`
}

type duration struct {
	Text string `json:"text"`
}
