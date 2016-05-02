package main

import (
	"encoding/json"
	"net/http"
)

const (
	ErrNotEnoughValues    = "not enough latitude and longitude values given"
	ErrInvalidCoordinates = "invalid latitude/longitude values"
	ErrGettingTimes       = "unable to get journey times"
)

func timesHandler(w http.ResponseWriter, r *http.Request) {
	//TODO this could go into the 'NewRefPoints' function?
	start := r.FormValue("start")
	end := r.FormValue("end")
	if start == "" || end == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Message: ErrNotEnoughValues,
		})
		return
	}
	journey, err := NewJourney(start, end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Message: ErrInvalidCoordinates,
		})
		return
	}

	err = journey.GetTimes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Message: ErrGettingTimes,
		})
		return
	}
	json.NewEncoder(w).Encode(journeyResponse{
		Times: journey.Times,
	})
}

type errorResponse struct {
	Message string `json:"message"`
}

type journeyResponse struct {
	Times journeyTimes `json:"journey_times"`
}
