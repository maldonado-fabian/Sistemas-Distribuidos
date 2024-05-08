package pdfapi

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func ProtectPDF(c *gin.Context) {
	fmt.Print("Ingrese la ruta del archivo PDF a proteger: ")
	var filePath string
	fmt.Scanln(&filePath)

	token := postAuth()
	server, taskID := getStartProtect()

	serverFilename := uploadFile(c, server, taskID, filePath, token)

	// Datos adicionales necesarios para processFile
	filename := filepath.Base(filePath) // o algún otro nombre lógico
	password := "userPassword"          // Debería venir de alguna parte segura

	processFile(c, server, taskID, serverFilename, filename, password, token)
}
