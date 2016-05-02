package main

import (
	"errors"
	"strconv"
	"strings"
)

type Journey struct {
	Start, End Coordinates
	Times      journeyTimes
}

type Coordinates struct {
	Latitude, Longitude float64
}

type journeyTimes struct {
	Car     string `json:"car"`
	Walk    string `json:"walk"`
	Bicycle string `json:"bicycle"`
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
	splitVals := strings.Split(vals, ",")

	c.Latitude, err = strconv.ParseFloat(splitVals[0], 64)
	if err != nil {
		return err
	}
	c.Longitude, err = strconv.ParseFloat(splitVals[1], 64)
	if err != nil {
		return err
	}

	if c.Valid() {
		return nil
	}
	return errors.New("invalid coordinates")
}

func (c Coordinates) Valid() bool {
	return (-90.0 <= c.Latitude && c.Latitude <= 90.0) &&
		(-180.0 <= c.Longitude && c.Longitude <= 180.0)
}

func (j *Journey) GetTimes(client GoogleMapsClient) error {
	var err error

	j.Times.Car, err = client.GetDuration(j.Start, j.End, "car")
	if err != nil {
		return err
	}

	j.Times.Walk, err = client.GetDuration(j.Start, j.End, "walking")
	if err != nil {
		return err
	}

	j.Times.Bicycle, err = client.GetDuration(j.Start, j.End, "bicycling")
	if err != nil {
		return err
	}

	return nil
}
