package main

import (
	"apidis/routes"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var opcion int
var id, nombre, apellido, rut, correo, contraseña string

// funcion para los que los inputs sean de la forma "Ingrese su nombre: "nombre" por alguna razon no acepta la primera vez que se llama
func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return input
}
func main() {
	// Establecer la conexión a la base de datos MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://fabu:izipizi123@distribuidos.wdrdmez.mongodb.net/?retryWrites=true&w=majority&appName=Distribuidos")
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

	// Manejar la entrada del usuario
	for opcion != 3 {
		fmt.Print("Bienvenido al sistema de protección de archivos de DiSis.\n" +
			"Para utilizar la aplicación seleccione los números\n" +
			"correspondientes al menú.\n")
		fmt.Println(" ")
		fmt.Println("Ingrese o registrese")
		fmt.Println(" ")
		fmt.Println("1) Ingreso")
		fmt.Println("2) Registro")
		fmt.Println("3) Salir")
		fmt.Println(" ")
		fmt.Scan(&opcion)
		fmt.Println(" ")
		switch opcion {
		case 1:
			correo = getInput("Ingrese su correo: ")
			contraseña = getInput("Ingrese su contraseña: ")
			fmt.Println(" ")
			// caso de login exitoso
			fmt.Println("ingreso exitoso!")
			fmt.Println(" ")
			for opcion != 3 {
				fmt.Print("menu principal\n")
				fmt.Println(" ")
				fmt.Println("1) clientes")
				fmt.Println("2) proteccion")
				fmt.Println("3) Salir")
				fmt.Println(" ")
				fmt.Scan(&opcion)
				fmt.Println(" ")
				switch opcion {
				case 1:
					for opcion != 7 {
						fmt.Println("1) Listar los clientes registrados")
						fmt.Println("2) Obtener un cliente por ID")
						fmt.Println("3) Obtener un cliente por RUT")
						fmt.Println("4) Registrar un nuevo cliente")
						fmt.Println("5) Actualizar datos de un cliente")
						fmt.Println("6) Borrar un cliente por ID")
						fmt.Println("7) Volver")
						fmt.Println(" ")
						fmt.Scan(&opcion)
						fmt.Println(" ")
						switch opcion {
						case 1:
							//routes.GetUsers
						case 2:
							//nose si esta bien
							id = getInput("Ingrese el ID del cliente: ")
							fmt.Println(" ")
							//routes.GetUserByID
						case 3:
							rut = getInput("Ingrese el RUT a buscar: ")
							fmt.Println(" ")
							//routes.GetUserByRUT
						case 4:
							nombre = getInput("Ingrese su nombre: ")
							apellido = getInput("Ingrese su apellido: ")
							rut = getInput("Ingrese su RUT: ")
							correo = getInput("Ingrese su correo: ")
							fmt.Println(" ")
							//routes.Register
							fmt.Println("¡Cliente", nombre, "creado con éxito!")
							fmt.Println(" ")
						case 5:
							nombre = getInput("Ingrese el nuevo nombre: ")
							apellido = getInput("Ingrese el nuevo apellido: ")
							rut = getInput("Ingrese el nuev	o RUT: ")
							correo = getInput("Ingrese el nuevo correo: ")
							fmt.Println(" ")
							//routes.UpdateUserByID
						case 6:
							id = getInput("Ingrese el ID del cliente a borrar: ")
							fmt.Println(" ")
							//routes.DeleteUserByID
						case 7:
							// vuelve a menu principal
						}
					}
				case 2:
					//wosh proteccion
				case 3:
					fmt.Println("¡vuelve pronto!")
				}
			}

		case 2:
			nombre = getInput("Ingrese su nombre: ")
			apellido = getInput("Ingrese su apellido: ")
			rut = getInput("Ingrese su RUT: ")
			correo = getInput("Ingrese su correo: ")
			contraseña = getInput("Ingrese su contraseña: ")
			fmt.Println(" ")
			// crear un nuevo usuario

			fmt.Println("¡registro exitoso!")
			fmt.Println(" ")
		case 3:
			fmt.Print("¡vuelve pronto!")
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
	// Iniciar el servidor HTTP
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
