package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// TODO this could be split into seperate functions
func TestTimesHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(timesHandler))
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

	if actualNoVals.Message != ErrNotEnoughValues {
		t.Errorf("expected %#v, got %#v", ErrNotEnoughValues, actualNoVals.Message)
	}

	// incorrect params
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

	// correct params
	query = "start=51.5034070,-0.1275920&end=51.4838940,-0.60440"
	url = server.URL + "/times?" + query

	resp = mustGetResponse(t, url)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status code: expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var actualValid journeyResponse

	err = json.NewDecoder(resp.Body).Decode(&actualValid)
	if err != nil {
		t.Fatalf("unexpected error when decoding json: got %s", err)
	}

	expected := journeyResponse{
		Times: journeyTimes{
			Car:     "48 mins",
			Bicycle: "2 hours 21 mins",
			Walk:    "7 hours 25 mins",
		},
	}

	if !reflect.DeepEqual(actualValid, expected) {
		t.Errorf("expected response to equal %#v, instead got %#v", expected, actualValid)
	}
}

func mustGetResponse(t *testing.T, url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
	return resp
}
