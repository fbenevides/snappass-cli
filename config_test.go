package main

import (
	"os"
	"testing"
)

func TestWriteAndReadConfig(t *testing.T) {
	// Create a temporary test file
	tempFile, err := os.CreateTemp("", "test_snappass_")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	defer func() {
		// Clean up after the test
		err := clear(tempFile.Name())
		if err != nil {
			t.Errorf("Error removing test file: %v", err)
		}
	}()

	// Test data
	testConfig := &Config{
		BaseUrl: "https://example.com",
	}

	// Write the test data to the test file
	err = write(tempFile.Name(), testConfig)
	if err != nil {
		t.Errorf("Error writing config: %v", err)
	}

	// Read the data back from the test file
	readConfig, err := read(tempFile.Name())
	if err != nil {
		t.Errorf("Error reading config: %v", err)
	}

	// Compare the read data with the original data
	if readConfig.BaseUrl != testConfig.BaseUrl {
		t.Errorf("Expected BaseUrl %s, got %s", testConfig.BaseUrl, readConfig.BaseUrl)
	}
}
