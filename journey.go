package main

import (
	"errors"
	"strconv"
	"strings"
)

type Journey struct {
	Start Coordinates
	End   Coordinates
}

type Coordinates struct {
	Latitude  float64
	Longitude float64
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
	if c.AreValid() {
		return nil
	}
	return errors.New("invalid coordinates")
}

func (c *Coordinates) AreValid() bool {
	return (-90.0 <= c.Latitude && c.Latitude <= 90.0) &&
		(-180.0 <= c.Longitude && c.Longitude <= 180.0)
}
