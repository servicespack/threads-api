package main

import (
	"github.com/servicespack/threads-api/controllers"
	"github.com/servicespack/threads-api/configs"
	"github.com/gin-gonic/gin"
)

type CreateThread struct {
	Text string `json:"text" binding:"required"`
}

func main() {
	configs.InitializeDatabase()

	r := gin.Default()
	r.GET("/threads", controllers.ListThreadas)
	r.POST("/threads", controllers.CreateThread)
	r.GET("/threads/:id", controllers.GetThread)
	r.PATCH("/threads/:id", controllers.UpdateThread)
	r.DELETE("/threads/:id", controllers.DeleteThread)

	r.Run()
}
