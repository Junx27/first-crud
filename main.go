package main

import (
	"github.com/Junx27/first-crud/config"
	"github.com/Junx27/first-crud/modules/user"
	"github.com/Junx27/first-crud/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	config.DBConnection()

	userRepository := repositories.NewUserRepository(config.DB)
	userService := user.NewUseCase(userRepository)
	userController := user.NewUserController(userService)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	user := router.Group("/user")
	user.GET("/all", userController.FindAll)

	router.Run(":8080")
}
