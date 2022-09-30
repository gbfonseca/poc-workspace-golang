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

	exists := projectExists(userPath)

	if !exists {

		err := createSources(userPath, currentPath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Setup finalizado com sucesso",
	})

}

func createSources(userProjPath string, currentPath string) error {
	if err := os.MkdirAll(userProjPath, os.ModeSticky|os.ModePerm); err != nil {

		return err
	}
	reactProjPath := path.Join(currentPath, "..", "react-proj/")

	utils.Copy(reactProjPath, userProjPath)
	return nil
}

func projectExists(userProjPath string) bool {
	_, err := os.ReadDir(userProjPath)

	if err != nil {
		return false
	}

	return true
}
