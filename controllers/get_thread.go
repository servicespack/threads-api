package controllers

import (
	"github.com/servicespack/threads-api/configs"
	"github.com/servicespack/threads-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetThread(context *gin.Context) {
	id := context.Param("id")

	db := configs.GetDB()
	var thread models.Thread
	err := db.First(&thread, id).Error
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, thread)
}
