package main

import (
	"apidis/models"
	. "apidis/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var opcion int
	var correo, contraseña string
	fmt.Println("Bienvenido al sistema de protección de archivos de DiSis.")

	for opcion != 3 {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1) Ingresar")
		fmt.Println("2) Registro")
		fmt.Println("3) Salir")

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
			// Luego de que el usuario inicie sesión con éxito, muestra las opciones del cliente
			if response.StatusCode == http.StatusOK {
				fmt.Println("Inicio de sesión exitoso. ¡Bienvenido!")
				menuPrincipal()
			} else {
				fmt.Println("Credenciales invalidas, te invitamos a registrarte")
			}
		case 2:
			// Logica para crear un nuevo usuario
			var nombre, apellido, rut, mail, pass string
			fmt.Println("Ingresa el nombre del nuevo usuario")
			fmt.Scan(&nombre)
			fmt.Println("Ingresa el apellido del nuevo usuario")
			fmt.Scan(&apellido)
			fmt.Println("Ingresa el rut del nuevo usuario")
			fmt.Scan(&rut)
			fmt.Println("Ingresa el mail del nuevo usuario")
			fmt.Scan(&mail)
			fmt.Println("Ingresa la contraseña del nuevo usuario")
			fmt.Scan(&pass)

			newUser := User{
				Name:     nombre,
				LastName: apellido,
				Rut:      rut,
				Email:    mail,
				Password: pass,
			}

			// Convertir el nuevo usuario a formato JSON
			newUserJSON, err := json.Marshal(newUser)
			if err != nil {
				fmt.Println("Error al convertir el usuario a JSON:", err)
				return
			}

			// Realizar la solicitud HTTP POST para crear el nuevo usuario
			response, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(newUserJSON))
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				return
			}
			defer response.Body.Close()

			// Verificar el código de estado de la respuesta
			if response.StatusCode != http.StatusCreated {
				fmt.Println("Error: El usuario no fue creado")
				return
			}

			fmt.Println("Usuario creado exitosamente")
		case 3:
			// Salir
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}

func menuCliente() {
	var opcionCliente int

	for opcionCliente != 6 {
		fmt.Println("Menu Clientes")
		fmt.Println("Seleccione una opción:")
		fmt.Println("1) Obtener todos los usuarios")
		fmt.Println("2) Obtener usuario por ID")
		fmt.Println("3) Eliminar usuario por ID")
		fmt.Println("4) Crear nuevo usuario")
		fmt.Println("5) Actualizar usuario")
		fmt.Println("6) Volver al menú anterior")

		fmt.Scan(&opcionCliente)

		switch opcionCliente {
		case 1:
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
				fmt.Println("--------------------")
				fmt.Printf("ID: %s\n Nombre: %s\n Apellido: %s\n Rut: %s\n Correo: %s\n", user.ID.Hex(), user.Name, user.LastName, user.Rut, user.Email)
				fmt.Println("--------------------")
			}
		case 2:
			// Lógica para obtener usuario por ID
			var idUsuario string
			fmt.Println("Ingresa el id del usuario que quieres econtrar")
			fmt.Scan(&idUsuario)
			response, err := http.Get("http://localhost:8080/users/" + idUsuario)
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				continue
			}
			defer response.Body.Close()

			// Verificar el código de estado de la respuesta
			if response.StatusCode != http.StatusOK {
				fmt.Println("Error: El usuario no fue encontrado")
				return
			}

			// Decodificar la respuesta JSON en un struct de usuario
			var user User
			if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
				fmt.Println("Error al decodificar la respuesta:", err)
				return
			}

			// Mostrar la información del usuario por pantalla
			fmt.Println("Usuario encontrado:")
			fmt.Println("--------------------")
			fmt.Printf("ID: %s\nNombre: %s\nApellido: %s\nRut: %s\nCorreo: %s\n", user.ID, user.Name, user.LastName, user.Rut, user.Email)
			fmt.Println("--------------------")

		case 3:
			// Lógica para eliminar usuario por ID
			// Realizar la solicitud HTTP DELETE para eliminar el usuario por ID
			var idUsuario string
			fmt.Println("Ingresa el id del usuario que quieres borrar")
			fmt.Scan(&idUsuario)
			req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/users/%s", idUsuario), nil)
			if err != nil {
				fmt.Println("Error al crear la solicitud HTTP:", err)
				return
			}

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				return
			}
			defer response.Body.Close()

			// Verificar el código de estado de la respuesta
			if response.StatusCode != http.StatusOK {
				fmt.Println("Error: El usuario no fue eliminado")
				return
			}

			fmt.Println("Usuario eliminado exitosamente")
		case 4:
			// Logica para crear un nuevo usuario
			var nombre, apellido, rut, mail, pass string
			fmt.Println("Ingresa el nombre del nuevo usuario")
			fmt.Scan(&nombre)
			fmt.Println("Ingresa el apellido del nuevo usuario")
			fmt.Scan(&apellido)
			fmt.Println("Ingresa el rut del nuevo usuario")
			fmt.Scan(&rut)
			fmt.Println("Ingresa el mail del nuevo usuario")
			fmt.Scan(&mail)
			fmt.Println("Ingresa la contraseña del nuevo usuario")
			fmt.Scan(&pass)

			newUser := User{
				Name:     nombre,
				LastName: apellido,
				Rut:      rut,
				Email:    mail,
				Password: pass,
			}

			// Convertir el nuevo usuario a formato JSON
			newUserJSON, err := json.Marshal(newUser)
			if err != nil {
				fmt.Println("Error al convertir el usuario a JSON:", err)
				return
			}

			// Realizar la solicitud HTTP POST para crear el nuevo usuario
			response, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewBuffer(newUserJSON))
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				return
			}
			defer response.Body.Close()

			// Verificar el código de estado de la respuesta
			if response.StatusCode != http.StatusCreated {
				fmt.Println("Error: El usuario no fue creado")
				return
			}

			fmt.Println("Usuario creado exitosamente")
		case 5:
			var id, nombre, apellido, rut, mail, pass string
			// ID del usuario que deseas actualizar
			fmt.Println("Ingresa el id actualizado del nuevo usuario")
			fmt.Scan(&id)
			// Scaneo de los datos
			fmt.Println("Ingresa el nombre actualizado del nuevo usuario")
			fmt.Scan(&nombre)
			fmt.Println("Ingresa el apellido actualizado del nuevo usuario")
			fmt.Scan(&apellido)
			fmt.Println("Ingresa el rut actualizado del nuevo usuario")
			fmt.Scan(&rut)
			fmt.Println("Ingresa el mail actualizado del nuevo usuario")
			fmt.Scan(&mail)
			fmt.Println("Ingresa la contraseña actualizado del nuevo usuario")
			fmt.Scan(&pass)

			// Datos actualizados del usuario
			updateData := User{
				Name:     nombre,
				LastName: apellido,
				Rut:      rut,
				Email:    mail,
				Password: pass,
			}

			// Convertir los datos de actualización a formato JSON
			updateDataJSON, err := json.Marshal(updateData)
			if err != nil {
				fmt.Println("Error al convertir los datos de actualización a JSON:", err)
				return
			}

			// Crear una nueva solicitud HTTP PUT para actualizar el usuario
			req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/users/%s", id), bytes.NewBuffer(updateDataJSON))
			if err != nil {
				fmt.Println("Error al crear la solicitud HTTP:", err)
				return
			}

			// Establecer el encabezado Content-Type
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				fmt.Println("Error al realizar la solicitud HTTP:", err)
				return
			}
			defer response.Body.Close()

			// Verificar el código de estado de la respuesta
			if response.StatusCode != http.StatusOK {
				fmt.Println("Error: El usuario no fue actualizado")
				return
			}

			fmt.Println("Usuario actualizado exitosamente")
		case 6:
			// Volver al menu principal
		default:
			fmt.Println("Opción no válida. Por favor, seleccione una opción válida.")
		}
	}
}
func menuPDF() {
	var id_protec, filePath, password string
	fmt.Println("Escriba el ID del cliente objetivo: ")
	fmt.Scan(&id_protec)
	fmt.Println("Escriba la ruta donde se encuentra el archivo (incluya el nombre): ")
	fmt.Scan(&filePath)

	response, err := http.Get("http://localhost:8080/users/" + id_protec)
	if err != nil {
		fmt.Println("Error al realizar la solicitud HTTP:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: El usuario no fue encontrado")
		return
	}

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		fmt.Println("Error al decodificar la respuesta:", err)
		return
	}
	password = user.Rut

	jsonData, _ := json.Marshal(map[string]string{"filePath": filePath, "password": password})
	response2, err2 := http.Post("http://localhost:8080/api/protect", "application/json", bytes.NewBuffer(jsonData))
	if err2 != nil {
		fmt.Println("Error al realizar la solicitud HTTP:", err2)
		return
	}
	defer response2.Body.Close()

	if response2.StatusCode == http.StatusOK {
		fmt.Println("¡Protección exitosa! Su archivo se encuentra en:", filePath)
	} else {
		fmt.Println("No se pudo proteger.")
	}
}

func menuPrincipal() {
	var opcion int
	fmt.Println("Menú Principal")
	for opcion != 3 {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1) Clientes")
		fmt.Println("2) Protección")
		fmt.Println("3) Salir")
		fmt.Scan(&opcion)
		switch opcion {
		case 1:
			menuCliente()
		case 2:
			menuPDF()
		case 3:
			// Retorna al menu anterior
		default:
			fmt.Println("ingresa una copcion valida")
		}
	}
}
