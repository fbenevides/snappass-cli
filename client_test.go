package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetPassword(t *testing.T) {
	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/set_password" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		// Simulate a successful response
		w.WriteHeader(http.StatusOK)
		response := Response{Status: http.StatusOK, Link: "example.com", Ttl: 3600}
		json.NewEncoder(w).Encode(response)
	}))
	defer ts.Close()

	// Create a client with the test server's URL
	client := NewClient(ts.URL)

	// Test a successful request
	response, err := client.SetPassword("test_password")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !response.IsSuccessful() {
		t.Error("expected successful response")
	}

	// Test a failed request
	ts.Close()
	response, err = client.SetPassword("test_password")
	if err == nil {
		t.Error("expected an error, got nil")
	}
	if response != nil {
		t.Errorf("expected nil response, got: %v", response)
	}
}

func TestSetPassword_HttpError(t *testing.T) {
	client := NewClient("http://nonexistent-url.com")

	// Test an error during HTTP request
	_, err := client.SetPassword("test_password")
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestResponse_IsSuccessful(t *testing.T) {
	tests := []struct {
		response *Response
		expected bool
	}{
		{&Response{Status: http.StatusOK}, true},
		{&Response{Status: http.StatusNotFound}, false},
		{&Response{Status: http.StatusInternalServerError}, false},
	}

	for _, test := range tests {
		if result := test.response.IsSuccessful(); result != test.expected {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}
