package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
}

func mustGetResponse(t *testing.T, url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
	return resp
}
