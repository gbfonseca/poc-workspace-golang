package controllers

import (
	"net/http"
	"os"
	"path"
	"workspace_go/main/utils"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Email string `form:"email" json:"email" xml:"email" binding:"required"`
}

func Setup(c *gin.Context) {

	var requestBody RequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	currentPath, err := utils.GetReactProj()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userPath, _ := utils.GetUserWorkspace(requestBody.Email)

	if err := os.MkdirAll(userPath, os.ModeSticky|os.ModePerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reactProjPath := path.Join(currentPath, "..", "react-proj/")

	utils.Copy(reactProjPath, userPath)

	c.JSON(200, gin.H{
		"message": "Setup finalizado com sucesso",
	})

}
