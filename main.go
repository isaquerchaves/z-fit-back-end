package main

import (
	"api/config"
	controllers "api/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDb()
}

func main() {
	app := gin.Default()

	app.GET("/muscles", controllers.GetAllMuscle)
	app.GET("/muscle/:id", controllers.GetMuscle)

	app.GET("/exercises", controllers.GetAllExercise)
	app.GET("exercise/:id", controllers.GetExercise)

	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}
