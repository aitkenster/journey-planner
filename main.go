package main

import "net/http"

const (
	ErrInvalidCoordinates = "invalid latitude/longitude values"
	ErrGettingTimes       = "unable to get journey times"
	ErrNoDataAvailable    = "not possible to calculate journey times between these coordinates"
)

func main() {
	http.HandleFunc("/times", timesHandler(Config{}))
	http.ListenAndServe(":8080", nil)
}

type Config struct {
	APIBaseURL string
}
