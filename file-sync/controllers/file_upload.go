package controllers

import (
	"net/http"
	"os/exec"
	"path/filepath"
	"workspace_go/main/utils"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	form, err := c.MultipartForm()

	email := form.Value["email"][0]
	file, err := c.FormFile("file")

	userWorkspacePath, err := utils.GetUserWorkspace(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := "miniapp" + extension

	userReactProj := userWorkspacePath + "/react-proj"

	c.SaveUploadedFile(file, userReactProj+"/"+newFileName)

	err = utils.Unzip(userReactProj+"/"+newFileName, userReactProj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("rm", "-rf", userReactProj+"/"+newFileName)
	go cmd.Run()

	// TODO - Transpile JSX -> JS

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload feito com sucesso",
	})

}
