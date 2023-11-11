package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteThread(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"updated": true,
	})
}
