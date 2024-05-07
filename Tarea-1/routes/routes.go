package routes

import (
	"context"
	"net/http"
	"strconv"

	"apidis/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Declaración de variable global para la colección de usuarios
var userCollection *mongo.Collection

// Función para configurar la colección de usuarios
func SetUserCollection(collection *mongo.Collection) {
	userCollection = collection
}

// Función para obtener todos los usuarios
func GetUsers(c *gin.Context) {
	var users []models.User

	cursor, err := userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuarios"})
		return
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &users)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al decodificar usuarios"})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

// Función para agregar un nuevo usuario
func PostUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	_, err := userCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar usuario"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

// Función para obtener un usuario por su ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	var user models.User
	err = userCollection.FindOne(context.Background(), bson.D{{"id", userID}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar usuario"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// Función para eliminar un usuario por su ID
func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	result, err := userCollection.DeleteOne(context.Background(), bson.D{{"id", userID}})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
		return
	}

	if result.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}

// Función para actualizar un usuario por su ID
func UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	result, err := userCollection.ReplaceOne(context.Background(), bson.D{{"id", userID}}, updatedUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
		return
	}

	if result.MatchedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

func Login(c *gin.Context) {
	var loginUser models.User
	if err := c.BindJSON(&loginUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	var existingUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": loginUser.Email}).Decode(&existingUser)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	if existingUser.Password != loginUser.Password {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Aquí generarías un token JWT
	// Por simplicidad, omitiré la generación del token en este ejemplo

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Inicio de sesión exitoso"})
}

// Ruta para el registro de usuario
func Register(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Datos de usuario inválidos"})
		return
	}

	// Verificar si el correo electrónico ya está en uso
	var existingUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": newUser.Email}).Decode(&existingUser)
	if err == nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "El correo electrónico ya está en uso"})
		return
	}

	// Insertar el nuevo usuario en la base de datos
	_, err = userCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar nuevo usuario"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Usuario registrado exitosamente"})
}

func Logout(c *gin.Context) {
	// Aquí podrías invalidar el token del usuario si es necesario
	// Por simplicidad, omitiré esta parte en este ejemplo

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Sesión cerrada exitosamente"})
}
