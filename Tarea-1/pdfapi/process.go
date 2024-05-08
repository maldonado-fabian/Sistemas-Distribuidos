package pdfapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type processResponse struct {
	Download_filename string `json:"download_filename"`
	Filesize          string `json:"filesize"`
	Output_filesize   string `json:"output_filesize"`
	Output_filenumber string `json:"output_filenumber"`
	Output_extensions string `json:"output_extensions"`
	Timer             string `json:"timer"`
	Status            string `json:"status"`
}

func processFile(c *gin.Context, server string, taskID string, serverFilename string, filename string, password string, token string) {
	// Crear el cuerpo de la solicitud
	reqBodyData := map[string]interface{}{
		"task": taskID,
		"tool": "protect",
		"files": []map[string]string{
			{
				"server_filename": serverFilename,
				"filename":        filename,
			},
		},
		"password": password,
	}
	reqBody, err := json.Marshal(reqBodyData)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal JSON"})
		return
	}

	// Construir la solicitud
	req, err := http.NewRequest("POST", "https://"+server+"/v1/process", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var responseJSON map[string]interface{}
	err = json.Unmarshal(respBody, &responseJSON)
	if err != nil {
		log.Printf("Error parsing response JSON: %v", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response JSON"})
		return
	}

	c.IndentedJSON(http.StatusOK, responseJSON)
}
