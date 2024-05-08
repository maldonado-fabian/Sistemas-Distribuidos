package pdfapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type startResponse struct {
	Server string `json:"server"`
	Task   string `json:"task"`
}

func getStartProtect() (string, string) {
	token := postAuth()

	req, _ := http.NewRequest("GET", "https://api.ilovepdf.com/v1/start/protect", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to process PDF: %v", err)
	}
	defer resp.Body.Close()

	var response startResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Failed to decode JSON: %v", err)
	}

	return response.Server, response.Task
}
