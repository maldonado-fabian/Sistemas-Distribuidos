package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `bson:"name"`
	Last_name string        `bson:"last_name"`
	Rut       string        `bson:"rut"`
	Email     string        `bson:"email"`
	Password  string        `bson:"password"`
}

func getMongoSession() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("urlmongo"))
	if err != nil {
		panic(err)
	}
	return session
}
func registerHandler(c *gin.Context) {
	fmt.Println("Ingrese su nombre:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Ingrese su apellido:")
	var lastName string
	fmt.Scanln(&lastName)

	fmt.Println("Ingrese su RUT:")
	var rut string
	fmt.Scanln(&rut)

	fmt.Println("Ingrese su correo:")
	var email string
	fmt.Scanln(&email)

	fmt.Println("Ingrese su contraseña:")
	var password string
	fmt.Scanln(&password)

	// TODO: Implement registration logic

	c.JSON(http.StatusOK, gin.H{"message": "¡Registro exitoso!"})
}

func clientsHandler(c *gin.Context) {
	fmt.Println("Menú clientes")
	fmt.Println("1) Listar los clientes registrados")
	fmt.Println("2) Obtener un cliente por ID")
	fmt.Println("3) Obtener un cliente por RUT")
	fmt.Println("4) Registrar un nuevo cliente")
	fmt.Println("5) Actualizar datos de un cliente")
	fmt.Println("6) Borrar un cliente por ID")
	fmt.Println("7) Volver")

	// TODO: Implement clients menu logic
}

func protectHandler(c *gin.Context) {
	fmt.Println("Escriba el ID del cliente objetivo:")
	var clientID string
	fmt.Scanln(&clientID)

	fmt.Println("Escriba la ruta donde se encuentra el archivo (incluya el nombre):")
	var filePath string
	fmt.Scanln(&filePath)

	fmt.Println("¡Protección exitosa! Su archivo se encuentra en:")
	fmt.Println(filePath)
}

func loginHandler(c *gin.Context) {
	session := getMongoSession()
	defer session.Close()

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var found User
	err := session.DB("test").C("users").Find(bson.M{"email": user.Email}).One(&found)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	if found.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	tokenString, _ := token.SignedString([]byte("secret"))
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", loginHandler)
		v1.POST("/register", registerHandler)
		v1.GET("/clients", clientsHandler)
		v1.POST("/protect", protectHandler)
	}

	router.Run(":" + os.Getenv("PORT"))
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Ingrese o regístrese")
		fmt.Println("1) Ingreso")
		fmt.Println("2) Registro")
		fmt.Println("3) Salir")
		fmt.Print("> ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			// Lógica para el ingreso
			fmt.Println("Ingreso seleccionado")
		case "2":
			// Lógica para el registro
			fmt.Println("Registro seleccionado")
		case "3":
			fmt.Println("¡Vuelve pronto!")
			os.Exit(0)
		default:
			fmt.Println("Opción no válida")
		}
	}
}
