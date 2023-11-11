package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListThreads(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"data": []interface{}{},
	})
}
