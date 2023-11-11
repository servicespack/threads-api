package controllers

import (
	"net/http"
	"github.com/servicespack/threads-api/configs"
	"github.com/servicespack/threads-api/models"
	"github.com/gin-gonic/gin"
)

type CreateThreadBody struct {
	Text string `json:"text" binding:"required"`
}

func CreateThread(context *gin.Context) {
	var createThreadBody CreateThreadBody
	if err := context.ShouldBind(&createThreadBody); err != nil {
		context.JSON(400, gin.H{"msg": err})
		return
	}

	db := configs.GetDB()
	db.Create(&models.Thread{
		Text: createThreadBody.Text,
	})

	context.JSON(http.StatusCreated, gin.H{
		"created": true,
	})
}
