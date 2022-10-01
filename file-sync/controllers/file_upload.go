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

	userPath, err := utils.GetUserWorkspace(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := "miniapp" + extension

	userReactProj := userPath + "/react-proj"

	c.SaveUploadedFile(file, userReactProj+"/"+newFileName)

	utils.Unzip(userPath+"/"+newFileName, userPath)

	cmd := exec.Command("yarn")
	cmd.Dir = userReactProj

	// Não Bloqueante
	// go cmd.Run()

	// Bloqueante
	cmd.Run()

	cmd = exec.Command("rm", "-rf", userReactProj+"/"+newFileName)
	go cmd.Run()

	// TODO - Transpile JSX -> JS

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload feito com sucesso",
	})

}
