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

	// test no params
	url := server.URL + "/times"

	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("incorrect status code: expected %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}

	expectedErr := errorResponse{
		Message: "not enough latitude and longitude values given",
	}

	var actualErr errorResponse
	err = json.NewDecoder(resp.Body).Decode(&actualErr)
	if err != nil {
		t.Errorf("unexpected error when decoding json: got %s", err)
	}

	if actualErr.Message != expectedErr.Message {
		t.Errorf("expected %#v, got %#v", expectedErr.Message, actualErr.Message)
	}
}
