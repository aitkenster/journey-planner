package main

import (
	"encoding/json"
	"net/http"
)

const (
	ErrNotEnoughValues    = "not enough latitude and longitude values given"
	ErrInvalidCoordinates = "invalid latitude/longitude values"
)

func main() {
	http.HandleFunc("/times", timesHandler)
	http.ListenAndServe(":8080", nil)
}

func timesHandler(w http.ResponseWriter, r *http.Request) {
	//TODO this could go into the 'NewRefPoints' function?
	start := r.FormValue("start")
	end := r.FormValue("end")
	if start == "" || end == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Message: ErrNotEnoughValues,
		})
	}
	_, err := NewJourney(start, end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Message: ErrInvalidCoordinates,
		})
	}
}

type errorResponse struct {
	Message string `json:"message"`
}
