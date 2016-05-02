package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestTimesHandlerInvalidParams(t *testing.T) {
	server := httptest.NewServer(timesHandler(Config{}))
	defer server.Close()

	// no params
	url := server.URL + "/times"

	resp := mustGetResponse(t, url)
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("incorrect status code: expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var actualNoVals errorResponse

	err := json.NewDecoder(resp.Body).Decode(&actualNoVals)
	if err != nil {
		t.Fatalf("unexpected error when decoding json: got %s", err)
	}

	if actualNoVals.Message != ErrInvalidCoordinates {
		t.Errorf("expected %#v, got %#v", ErrInvalidCoordinates, actualNoVals.Message)
	}

	// invalid values
	query := "start=51.5034070,-0.1275920&end=51.4838940,-0.60440ABC"
	url = server.URL + "/times?" + query

	resp = mustGetResponse(t, url)
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("incorrect status code: expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	var actualIncorrectVals errorResponse

	err = json.NewDecoder(resp.Body).Decode(&actualIncorrectVals)
	if err != nil {
		t.Fatalf("unexpected error when decoding json: got %s", err)
	}

	if actualIncorrectVals.Message != ErrInvalidCoordinates {
		t.Errorf("expected %#v, got %#v", ErrInvalidCoordinates, actualIncorrectVals.Message)
	}
}

func TestTimesHandlerValidParams(t *testing.T) {
	minutes := 5
	googleMapsServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(newGoogleMapsResponse(minutes))
			minutes = minutes + 5
		}),
	)
	defer googleMapsServer.Close()

	server := httptest.NewServer(timesHandler(Config{
		APIBaseURL: googleMapsServer.URL,
	}))
	defer server.Close()

	query := "start=51.5034070,-0.1275920&end=51.4838940,-0.60440"
	url := server.URL + "/times?" + query

	resp := mustGetResponse(t, url)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status code: expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var actual journeyResponse

	err := json.NewDecoder(resp.Body).Decode(&actual)
	if err != nil {
		t.Fatalf("unexpected error when decoding json: got %s", err)
	}

	expected := journeyResponse{
		Times: journeyTimes{
			Car:     "5 mins",
			Walk:    "10 mins",
			Bicycle: "15 mins",
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected response to equal %#v, instead got %#v", expected, actual)
	}
}

func TestTimesHandlerNoResults(t *testing.T) {
	googleMapsServer := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(GoogleMapsResponse{
				Rows: []row{{Elements: []*element{{Status: "ZERO_RESULTS"}}}},
			})
		}),
	)
	defer googleMapsServer.Close()

	server := httptest.NewServer(timesHandler(Config{
		APIBaseURL: googleMapsServer.URL,
	}))
	defer server.Close()

	query := "start=51.5034070,-0.1275920&end=51.4838940,-70.60440"
	url := server.URL + "/times?" + query

	resp := mustGetResponse(t, url)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status code: expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var actual errorResponse

	err := json.NewDecoder(resp.Body).Decode(&actual)
	if err != nil {
		t.Fatalf("unexpected error when decoding json: got %s", err)
	}

	if actual.Message != ErrNoDataAvailable {
		t.Errorf("expected %#v, got %#v", ErrNoDataAvailable, actual.Message)
	}

}

func mustGetResponse(t *testing.T, url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
	return resp
}

func newGoogleMapsResponse(mins int) GoogleMapsResponse {
	durationText := fmt.Sprintf("%v mins", mins)
	return GoogleMapsResponse{
		Rows: []row{
			{Elements: []*element{{
				Duration: duration{Text: durationText},
				Status:   "OK",
			}}},
		},
	}
}
