package main

import (
	"workspace_go/main/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"192.168.15.113"})

	v1 := router.Group("/v1/workspace")
	{
		v1.GET("/health", controllers.Health)
		v1.POST("/setup", controllers.Setup)
		v1.POST("/uploadAll", controllers.UploadAll)

	}

	router.Run("localhost:3000")
}
