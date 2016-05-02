package main

import (
	"encoding/json"
	"net/http"
)

const (
	ErrNotEnoughValues = "not enough latitude and longitude values given"
)

func main() {
	http.HandleFunc("/times", timesHandler)
	http.ListenAndServe(":8080", nil)
}

func timesHandler(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	if start == "" || end == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{
			Message: ErrNotEnoughValues,
		})
	}
}

type errorResponse struct {
	Message string `json:"message"`
}
