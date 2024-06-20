package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func UpdateThread(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"updated": true,
	})
}
