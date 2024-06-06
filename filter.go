package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

// TestResult represents the structure of each test result object.
type TestResult struct {
	Name         string `json:"name"`
	DurationMS   int    `json:"duration_ms"`
	TotalTests   int    `json:"total_tests"`
	FailedTests  int    `json:"failed_tests"`
	SkippedTests int    `json:"skipped_tests"`
	PassedTests  int    `json:"passed_tests"`
	FailPct      int    `json:"fail_pct"`
}

// JSONData represents the overall structure of the JSON file.
type JSONData struct {
	Data struct {
		TotalPages    int `json:"totalPages"`
		TotalItems    int `json:"totalItems"`
		PageItemCount int `json:"pageItemCount"`
		PageSize      int `json:"pageSize"`
	} `json:"data"`
	Content []TestResult `json:"content"`
}

func main() {
	// Define a flag for the time stamp
	timeStampMinutes := flag.Int("t", 0, "time stamp in minutes")
	flag.Parse()

	if *timeStampMinutes <= 0 {
		fmt.Println("Please provide a valid time stamp in minutes using the -t flag.")
		return
	}

	// Convert time stamp to milliseconds
	timeStampMS := *timeStampMinutes * 60 * 1000

	// Open the JSON file
	file, err := os.Open("testData.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Read the JSON file
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal the JSON data into JSONData struct
	var jsonData JSONData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return
	}

	// Filter the results based on the duration_ms
	var filteredResults []TestResult
	for _, result := range jsonData.Content {
		if result.DurationMS > timeStampMS {
			filteredResults = append(filteredResults, result)
		}
	}

	// Print the filtered results
	fmt.Println("Filtered Results:")
	for _, result := range filteredResults {
		fmt.Printf("%+v\n", result)
	}
}
