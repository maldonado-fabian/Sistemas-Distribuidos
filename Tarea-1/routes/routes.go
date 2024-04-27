package routes

import (
	"net/http"

	. "apidis/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []User{
	{Id: 1, Name: "fabian", LastName: "Maldonado", Rut: "12345678 - 1", Email: "fm@gmail.com", Password: "izipizi123"},
	{Id: 2, Name: "rodo", LastName: "Osorio", Rut: "12345678 - 2", Email: "ro@gmail.com", Password: "izipizi123"},
	{Id: 3, Name: "franco", LastName: "Nose", Rut: "12345678 - 3", Email: "fn@gmail.com", Password: "izipizi123"},
}

// getAlbums responds with the list of all albums as JSON.
func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	var newUser User

	c.BindJSON(&newUser)

	// Add the new album to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	idint, _ := strconv.Atoi(id)

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.Id == idint {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)

	// Buscar el índice del usuario en la lista
	index := -1
	for i, user := range users {
		if user.Id == idint {
			index = i
			break
		}
	}

	// Si se encontró el usuario, eliminarlo de la lista
	if index != -1 {
		users = append(users[:index], users[index+1:]...)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
	}
}

func UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	idint, _ := strconv.Atoi(id)

	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Buscar el índice del usuario en la lista
	index := -1
	for i, user := range users {
		if user.Id == idint {
			index = i
			break
		}
	}

	// Si se encontró el usuario, actualizar sus datos
	if index != -1 {
		// Conservar los valores existentes de los campos no proporcionados
		if updatedUser.Name == "" {
			updatedUser.Name = users[index].Name
		}
		if updatedUser.LastName == "" {
			updatedUser.LastName = users[index].LastName
		}
		if updatedUser.Rut == "" {
			updatedUser.Rut = users[index].Rut
		}
		if updatedUser.Email == "" {
			updatedUser.Email = users[index].Email
		}
		if updatedUser.Password == "" {
			updatedUser.Password = users[index].Password
		}

		// Asignar el ID existente al usuario actualizado
		updatedUser.Id = users[index].Id

		users[index] = updatedUser

		c.IndentedJSON(http.StatusOK, users[index])
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
	}
}
