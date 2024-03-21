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

	app.GET("/muscles", Muscle.GetAllMuscle)
	app.GET("/muscle/:id", Muscle.GetMuscle)

	app.Run("localhost:3001")
}
