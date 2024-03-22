package controllers

import (
	"api/config"
	"api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	// Criar uma nova sessão para garantir que o filtro não afete outras consultas
	session := config.DB.Session(&gorm.Session{})

	var muscle models.Muscle
	session.First(&muscle, "id = ?", id)

	c.JSON(200, gin.H{
		"muscle": muscle,
	})
}
