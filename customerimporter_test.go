package main

import (
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestProcessCustomerData(t *testing.T) {
	// Create a temporary CSV file for testing
	tempFile, err := os.CreateTemp("", "test_customers.csv")
	if err != nil {
		t.Fatal("Error creating temporary file:", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Write test data to the temporary CSV file
	testData := [][]string{
		{"first_name", "last_name", "email", "gender", "ip_address"},
		{"Mildred", "Hernandez", "mhernandez0@github.io", "Female", "38.194.51.128"},
		{"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
		{"Dennis", "Henry", "dhenry2@hubpages.com", "Male", "155.75.186.217"},
		{"InvalidRowWithoutEmail", "", "", "", ""}, // Add empty fields to match the structure
	}

	writer := csv.NewWriter(tempFile)
	for _, row := range testData {
		if err := writer.Write(row); err != nil {
			t.Fatal("Error writing to temporary file:", err)
		}
	}
	writer.Flush()

	// Call the function under test
	result := processCustomerData(tempFile.Name())

	// Define the expected result based on the test data
	expectedResult := map[string]int{
		"github.io":       1,
		"cyberchimps.com": 1,
		"hubpages.com":    1,
		// Add more expected results based on test data
	}

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Unexpected result. Expected: %v, Got: %v", expectedResult, result)
	}
}
