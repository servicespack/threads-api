package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/threads", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": []interface{}{},
		})
	})
	r.Run()
}
