package controllers

import (
	"api/config"
	"api/models"

	"github.com/gin-gonic/gin"
)

func GetAllExercise(c *gin.Context) {
	var exercise []models.Exercise
	config.DB.Find(&exercise)

	c.JSON(200, gin.H{
		"exercise": exercise,
	})
}
