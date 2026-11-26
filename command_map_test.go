package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/dweatherstone/pokedexcli/internal/pokeapi"
)

func TestCommandMap(t *testing.T) {
	testMapCommand(t, "commandMap", commandMap, "/next-page")
}

func TestCommandMapB(t *testing.T) {
	testMapCommand(t, "commandMapB", commandMapB, "/previous-page")
}

func TestCommandMapInitialPage(t *testing.T) {
	testMapCommand(t, "commandMap", commandMap, "/")
}

func TestCommandMapBOnFirstPage(t *testing.T) {
	cfg := &config{
		pokeClient: pokeapi.NewClient(),
		mapPrevURL: nil,
	}

	output := captureStdout(t, func() {
		if err := commandMapB(cfg, []string{}); err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})

	expectedOutput := "You're on the first page\n"
	if output != expectedOutput {
		t.Errorf("expected output %q, got %q", expectedOutput, output)
	}
}

func testMapCommand(t *testing.T, name string, command func(*config, []string) error, expectedPath string) {
	t.Helper()

	// Create a mock API server.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Errorf("expected request path %q, got %q", expectedPath, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := io.WriteString(w, `{
		"count": 2,
		"next": "https://example.com/next",
		"previous": "https://example.com/previous",
		"results": [
			{"name": "canalave-city-area", "url": "https://example.com/1/"},
			{"name": "eterna-city-area", "url": "https://example.com/2/"}
		]
		}`)
		if err != nil {
			t.Errorf("failed to write response: %v", err)
		}
	}))
	defer server.Close()

	// Set up config with the mock server URLs
	cfg := &config{
		pokeClient:      pokeapi.NewClient(),
		locationAreaURL: server.URL,
	}

	if expectedPath != "/" {
		cfg.mapNextURL = stringPointer(server.URL + expectedPath)
		cfg.mapPrevURL = stringPointer(server.URL + "/previous-page")
	}

	// Capture stdout so we can check the printed location names.
	output := captureStdout(t, func() {
		if err := command(cfg, []string{}); err != nil {
			t.Fatalf("%s returned an unexpected error: %v", name, err)
		}
	})

	expectedOutput := "canalave-city-area\neterna-city-area\n"
	if output != expectedOutput {
		t.Errorf("unexpected output:\nexpected: %q\ngot:       %q", expectedOutput, output)
	}

	// Check that the pagination URLs were updated
	if cfg.mapNextURL == nil || *cfg.mapNextURL != "https://example.com/next" {
		t.Errorf("unexpected mapNextURL: %v", cfg.mapNextURL)
	}
	if cfg.mapPrevURL == nil || *cfg.mapPrevURL != "https://example.com/previous" {
		t.Errorf("unexpected mapPrevURL: %v", cfg.mapPrevURL)
	}
}

func stringPointer(s string) *string {
	return &s
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdout pipe: %v", err)
	}

	os.Stdout = w

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close stdout writer: %v", err)
	}
	os.Stdout = oldStdout

	output, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("failed to read stdout: %v", err)
	}

	if err := r.Close(); err != nil {
		t.Fatalf("failed to close stdout reader: %v", err)
	}
	return string(output)
}
