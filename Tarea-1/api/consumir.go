package main

import (
	"apidis/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var opcion int
	var correo, contraseña string

	// Ciclo principal del programa
	for opcion != 6 {
		fmt.Println("Bienvenido al sistema de protección de archivos de DiSis.")
		fmt.Println("Seleccione una opción:")
		fmt.Println("1) Ingresar")
		fmt.Println("2) Registro")
		fmt.Println("3) Obtener todos los usuarios")
		fmt.Println("4) Obtener usuario por ID")
		fmt.Println("5) Eliminar usuario por ID")
		fmt.Println("6) Salir")

		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			// Ingresar
			fmt.Println("Ingrese su correo:")
			fmt.Scan(&correo)
			fmt.Println("Ingrese su contraseña:")
			fmt.Scan(&contraseña)

			// Llamar a la función de inicio de sesión de la API
			response, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer([]byte(`{"email":"`+correo+`","password":"`+contraseña+`"}`)))
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				continue
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				fmt.Println("Inicio de sesión exitoso. ¡Bienvenido!")
				// Aquí iría la lógica del CRUD
				// Por ejemplo, puedes llamar a una función que maneje las operaciones de CRUD
			} else {
				fmt.Println("Inicio de sesión fallido. Verifique su correo y contraseña.")
			}
		case 2:
			// Registro
			fmt.Println("Lógica de registro aún no implementada.")
		case 3:
			// Obtener todos los usuarios
			response, err := http.Get("http://localhost:8080/users")
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				continue
			}
			defer response.Body.Close()

			var users []models.User
			if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
				fmt.Println("Error al decodificar la respuesta:", err)
				continue
			}

			fmt.Println("Usuarios:")
			for _, user := range users {
				fmt.Printf("ID: %s, Nombre: %s, Apellido: %s, Rut: %s, Correo: %s\n", user.ID.Hex(), user.Name, user.LastName, user.Rut, user.Email)
			}
		case 4:
			// Obtener usuario por ID
			var id string
			fmt.Println("Ingrese el ID del usuario:")
			fmt.Scan(&id)

			response, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%s", id))
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				continue
			}
			defer response.Body.Close()

			var user models.User
			if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
				fmt.Println("Error al decodificar la respuesta:", err)
				continue
			}

			fmt.Printf("Usuario encontrado - ID: %s, Nombre: %s, Apellido: %s, Rut: %s, Correo: %s\n", user.ID.Hex(), user.Name, user.LastName, user.Rut, user.Email)
		case 5:
			// Eliminar usuario por ID
			var id string
			fmt.Println("Ingrese el ID del usuario que desea eliminar:")
			fmt.Scan(&id)

			req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/users/%s", id), nil)
			if err != nil {
				fmt.Println("Error al crear la solicitud HTTP:", err)
				continue
			}

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				continue
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				fmt.Println("Usuario eliminado correctamente.")
			} else {
				fmt.Println("Error al eliminar el usuario.")
			}
		case 6:
			// Salir
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}
