package main

import (
	"encoding/json"
	"net/http"
)

func timesHandler(config Config) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		journey, err := NewJourney(r.FormValue("start"), r.FormValue("end"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorResponse{
				Message: ErrInvalidCoordinates,
			})
			return
		}

		client := NewGoogleMapsClient(config)

		err = journey.GetTimes(client)
		if err != nil {
			errMsg := err.Error()
			if errMsg != ErrNoDataAvailable {
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(errorResponse{
				Message: errMsg,
			})
			return
		}

		json.NewEncoder(w).Encode(journeyResponse{
			Times: journey.Times,
		})
	})
}

type errorResponse struct {
	Message string `json:"message"`
}

type journeyResponse struct {
	Times journeyTimes `json:"journey_times"`
}
