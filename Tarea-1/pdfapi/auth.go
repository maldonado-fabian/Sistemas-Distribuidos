package pdfapi

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const public_key string = "project_public_c14b91af1ac4799b1f229e682d664154_x5xfs2fff307d4bde7d8b8b6bef3a0f2a9391"

type authResponse struct {
	Token string `json:"token"`
}

func postAuth() string {
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
