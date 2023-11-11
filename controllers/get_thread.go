package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetThread(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"get": true,
	})
}
