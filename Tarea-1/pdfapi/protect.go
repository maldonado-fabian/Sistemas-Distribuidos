package pdfapi

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ProtectPDF(c *gin.Context) {
	var input struct {
		FilePath string `json:"filePath"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token := postAuth()
	server, taskID := getStartProtect()

	serverFilename := uploadFile(c, server, taskID, input.FilePath, token)

	// Datos adicionales necesarios para processFile
	filename := filepath.Base(input.FilePath) // o algún otro nombre lógico

	processFile(c, server, taskID, serverFilename, filename, input.Password, token)
}
