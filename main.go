package main

import (
	"api/config"
	controllers "api/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
}

func main() {
	app := gin.Default()

	// Middleware CORS
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	app.GET("/muscles", controllers.GetAllMuscle)
	app.GET("/muscle/:id", controllers.GetMuscle)

	app.GET("/exercises", controllers.GetAllExercise)
	app.GET("/exercise/id/:id", controllers.GetExercise)
	app.GET("/exercise/muscleid/:muscle_id", controllers.GetExerciseMuscleId)

	app.Run("localhost:3004")
}
