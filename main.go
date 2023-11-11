package main

import (
	"github.com/gin-gonic/gin"
	"github.com/servicespack/threads-api/configs"
	"github.com/servicespack/threads-api/models"
	"net/http"
)

type CreateThread struct {
	Text string `json:"text" binding:"required"`
}

func main() {
	configs.InitializeDatabase()

	r := gin.Default()
	r.GET("/threads", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": []interface{}{},
		})
	})

	r.POST("/threads", func(context *gin.Context) {
		var createThread CreateThread
		if err := context.ShouldBind(&createThread); err != nil {
			context.JSON(400, gin.H{"msg": err})
			return
		}

		db := configs.GetDB()
		db.Create(&models.Thread{
            Text: createThread.Text,
        })

		context.JSON(http.StatusCreated, gin.H{
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
