package pokeapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"count": 1541, "next": "https://example.com/next", "previous": null, "results": [{"name": "canalave-city-area", "url": "https://example.com/1/"}]}`))
	}))
	defer server.Close()

	client := NewClient()

	var response LocationAreaResponse
	err := client.Get(server.URL, &response)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response.Count != 1541 {
		t.Fatalf("expected count 1541, got %d", response.Count)
	}
	if response.Next == nil {
		t.Fatal("expected next URL, got nil")
	}
	if *response.Next != "https://example.com/next" {
		t.Errorf("unexpected next URL: %s", *response.Next)
	}
	if response.Previous != nil {
		t.Errorf("expected previous to be nil, got %v", *response.Previous)
	}
	if len(response.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(response.Results))
	}
	if response.Results[0].Name != "canalave-city-area" {
		t.Errorf("unexpected location area name: %s", response.Results[0].Name)
	}
}

func TestClientGetHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
	}))
	defer server.Close()

	client := NewClient()

	var response LocationAreaResponse
	err := client.Get(server.URL, &response)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestClientGetInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("this isn't valid JSON"))
	}))
	defer server.Close()

	client := NewClient()

	var response LocationAreaResponse
	err := client.Get(server.URL, &response)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestClientGetRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.Close()

	client := NewClient()

	var response LocationAreaResponse
	err := client.Get(server.URL, &response)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
