package main

import (
	"apidis/pdfapi"
	"apidis/routes"
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection

func main() {
	// Establecer la conexión a la base de datos MongoDB
	envFile := filepath.Join("..", ".env")
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtiene la URL de conexión a la base de datos desde las variables de entorno
	dbURL := os.Getenv("MONGO_URI")
	port := os.Getenv("PORT")

	clientOptions := options.Client().ApplyURI(dbURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Verificar la conexión a la base de datos
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Seleccionar la colección de usuarios
	userCollection = client.Database("t1distribuidos").Collection("users")

	// Configurar la colección de usuarios en el paquete routes
	routes.SetUserCollection(userCollection)

	// Inicializar el enrutador Gin
	router := gin.Default()

	// Asignar las rutas a las funciones de manejo definidas en el paquete routes
	router.GET("/users", routes.GetUsers)
	router.POST("/users", routes.PostUser)
	router.GET("/users/:id", routes.GetUserByID)
	router.DELETE("/users/:id", routes.DeleteUserByID)
	router.PUT("/users/:id", routes.UpdateUserByID)
	router.POST("/login", routes.Login)
	router.POST("/register", routes.Register)
	router.POST("/logout", routes.Logout)
	router.POST("/api/protect", pdfapi.ProtectPDF)

	// Iniciar el servidor HTTP
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
