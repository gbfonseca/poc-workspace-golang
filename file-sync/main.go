package main

import (
	"net/http"
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
		v1.POST("/fileupload", controllers.FileUpload)
		v1.StaticFS("/storage", http.Dir("../storage"))
		v1.POST("/share", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{})
		})

	}

	router.Run("localhost:3000")
}
