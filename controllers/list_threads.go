package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListThreadas(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"data": []interface{}{},
	})
}
