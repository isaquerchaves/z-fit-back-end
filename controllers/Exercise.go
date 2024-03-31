package controllers

import (
	"api/config"
	"api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllExercise(c *gin.Context) {
	var exercise []models.Exercise
	config.DB.Find(&exercise)

	c.JSON(200, gin.H{
		"exercise": exercise,
	})
}

func GetExercise(c *gin.Context) {
	id := c.Param("id")

	session := config.DB.Session(&gorm.Session{})

	var exercise models.Exercise
	session.First(&exercise, "id = ?", id)

	c.JSON(200, gin.H{
		"exercise": exercise,
	})
}

func GetExerciseMuscleId(c *gin.Context) {
	muscle_id := c.Param("muscle_id")

	session := config.DB.Session(&gorm.Session{})

	var exercise []models.Exercise
	session.Find(&exercise, "muscle_id = ?", muscle_id)

	c.JSON(200, gin.H{
		"exercise": exercise,
	})
}
