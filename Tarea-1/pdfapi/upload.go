package pdfapi

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type uploadResponse struct {
	Server_filename string `json:"server_filename"`
}

func uploadFile(c *gin.Context, server string, taskID string, filePath string, token string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.WriteField("task", taskID)
	writer.Close()

	req, _ := http.NewRequest("POST", "https://"+server+"/v1/upload", body)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var uploadResp uploadResponse
	json.NewDecoder(resp.Body).Decode(&uploadResp)

	return uploadResp.Server_filename
}
