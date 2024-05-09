package pdfapi

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type authResponse struct {
	Token string `json:"token"`
}

func postAuth() string {
	envFile := filepath.Join("..", ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtiene la URL de conexión a la base de datos desde las variables de entorno
	public_key := os.Getenv("PUBLIC_KEY")

	body, _ := json.Marshal(map[string]string{
		"public_key": public_key,
	})

	req, _ := http.NewRequest("POST", "https://api.ilovepdf.com/v1/auth", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error al obtener token: %v", err)
	}
	defer resp.Body.Close()

	var res authResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Fatalf("Error al decodificar respuesta: %v", err)
	}

	return res.Token
}
