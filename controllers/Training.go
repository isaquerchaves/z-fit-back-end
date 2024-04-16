package controllers

import (
	"api/config"
	"api/models"

	"github.com/gin-gonic/gin"
)

func GetTrainingSplit(c *gin.Context) {
	var split []models.TrainingSplit
	config.DB.Find(&split)

	c.JSON(200, gin.H{
		"split": split,
	})
}
