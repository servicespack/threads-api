package main

import (
	"github.com/gin-gonic/gin"
	"github.com/servicespack/threads-api/configs"
	"net/http"
)

func main() {
	configs.InitializeDatabase()

	r := gin.Default()
	r.GET("/threads", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": []interface{}{},
		})
	})
	r.Run()
}
