package controllers

import (
	"net/http"
	"os/exec"
	"path"
	"path/filepath"
	"workspace_go/main/lib"
	"workspace_go/main/utils"

	"github.com/gin-gonic/gin"
)

func UploadAll(c *gin.Context) {
	form, err := c.MultipartForm()

	email := form.Value["email"][0]
	file, _ := c.FormFile("file")

	userWorkspacePath, err := utils.GetUserWorkspace(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := "miniapp" + extension

	userProj := path.Join(userWorkspacePath, "user-proj", "src")

	c.SaveUploadedFile(file, userProj+"/"+newFileName)

	err = utils.Unzip(userProj+"/"+newFileName, userProj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cmd := exec.Command("rm", "-rf", userProj+"/"+newFileName)
	go cmd.Run()

	// TODO - Transpile JSX -> JS
	err = lib.PostFiles(userProj, userWorkspacePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload feito com sucesso",
	})

}
