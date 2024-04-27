package main

import (
	. "apidis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", GetUsers)
	router.POST("/users", PostUser)
	router.GET("/users/:id", GetUserByID)
	router.DELETE("/users/:id", DeleteUserByID)
	router.PUT("/users/:id", UpdateUserByID)
	router.Run("localhost:8080")
}
