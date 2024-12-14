package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Data string `json:"data"`
}

func main() {
	urlPtr := flag.String("url", "", "API endpoint URL")
	flag.Parse()

	if *urlPtr == "" {
		log.Fatal("'--url' flag is required.")
	}

	requestBody := RequestBody{}
	// Read from stdin
	bodyBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %v", err)
	}
	requestBody.Data = string(bodyBytes)


	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	resp, err := http.Post(*urlPtr, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated { //Handle various success codes
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Fatalf("POST request failed with status code %d: %s", resp.StatusCode, string(bodyBytes))
	}

	fmt.Printf("POST request successful. Status code: %d\n", resp.StatusCode)
}
