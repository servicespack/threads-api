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

	r.POST("/threads", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"created": true,
		})
	})

	r.GET("/threads/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"get": true,
		})
	})

	r.PATCH("/threads/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"updated": true,
		})
	})

	r.DELETE("/threads/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"updated": true,
		})
	})

	r.Run()
}
