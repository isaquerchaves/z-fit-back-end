package main

import (
	"api/config"
	Muscle "api/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
}

func main() {
	app := gin.Default()

	app.GET("/muscle", Muscle.GetAllMuscle)

	app.Run("localhost:3001")
}
