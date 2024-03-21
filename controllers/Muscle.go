package Muscle

import (
	"api/config"
	"api/models"

	"github.com/gin-gonic/gin"
)

func GetAllMuscle(c *gin.Context) {
	var muscle []models.Muscle
	config.DB.Find(&muscle)

	c.JSON(200, gin.H{
		"muscles": muscle,
	})
}
