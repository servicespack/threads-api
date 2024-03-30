package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/servicespack/threads-api/configs"
	"github.com/servicespack/threads-api/models"
)

func DeleteThread(context *gin.Context) {
	id := context.Param("id")
	db := configs.GetDB()

	err := db.Delete(&models.Thread{}, id)
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"deleted": true,
	})
}
