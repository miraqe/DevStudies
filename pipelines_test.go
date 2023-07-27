package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetDealsHandler(t *testing.T) {
	// Load the config before running the test
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}

	// Set the API token as an environment variable
	os.Setenv("PIPEDRIVE_API_TOKEN", config.PipedriveAPIToken)

	// Create a mock-up http.ResponseWriter
	w := httptest.NewRecorder()

	// Create a mock-up http.Request for the GET request
	r := httptest.NewRequest(http.MethodGet, "/getDeals", nil)

	// Call getDealsHandler with the mock-up request and response
	getDealsHandler(w, r)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check if the response body contains the "title" field
	expectedTitle := "Ingretchen"
	if !strings.Contains(w.Body.String(), expectedTitle) {
		t.Errorf("Expected title '%s' not found in the response", expectedTitle)
	}
}

func TestAddDealHandler(t *testing.T) {
	// Load the config before running the test
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}

	// Set the API token as an environment variable
	os.Setenv("PIPEDRIVE_API_TOKEN", config.PipedriveAPIToken)

	// Prepare the payload data for the new deal
	payloadData := map[string]interface{}{
		"title":              "Craddle Merch",
		"value":              576,
		"currency":           "EUR",
		"status":             "open",
		"org_id":             1,
		"participants_count": 1,
	}

	// Convert the payload data to JSON format
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		t.Fatalf("Error converting payload to JSON: %s", err)
	}

	// Create a new POST request with the payload
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/addDeal", bytes.NewBuffer(payloadBytes))
	r.Header.Set("Content-Type", "application/json")

	// Call addDealHandler with the mock-up request, response, and payload data
	addDealHandler(w, r)

	// Check the response status code
	log.Println()
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, w.Code)
	}
}

func TestChangeDealHandler(t *testing.T) {
	// Load the config before running the test
	config, err := loadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}

	// Set the API token as an environment variable
	os.Setenv("PIPEDRIVE_API_TOKEN", config.PipedriveAPIToken)

	// Prepare the payload data for changing deal 44
	payloadData := map[string]interface{}{
		"title": "Spruce Bravo",
		"value": 2634,
	}

	// Convert the payload data to JSON format
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		t.Fatalf("Error converting payload to JSON: %s", err)
	}

	// Create a new PUT request with the payload to change deal 44
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPut, "/changeDeal", bytes.NewBuffer(payloadBytes))
	r.Header.Set("Content-Type", "application/json")

	// Call changeDealHandler with the mock-up request, response, and payload data
	changeDealHandler(w, r)

	// Check the response status code
	log.Println()
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}
