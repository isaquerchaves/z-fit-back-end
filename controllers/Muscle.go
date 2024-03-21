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

func GetMuscle(c *gin.Context) {
	id := c.Param("id")

	var muscle models.Muscle
	config.DB.First(&muscle, "id = ?", id)

	c.JSON(200, gin.H{
		"muscle": muscle,
	})
}
