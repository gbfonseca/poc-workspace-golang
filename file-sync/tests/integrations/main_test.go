package integrations_test

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}