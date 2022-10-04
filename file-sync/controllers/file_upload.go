package controllers

import (
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"workspace_go/main/lib"
	"workspace_go/main/utils"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	form, _ := c.MultipartForm()

	email := form.Value["email"][0]
	filePath := form.Value["filepath"][0]
	// why := form.Value["why"][0]
	// miniAppVersion := form.Value["miniAppVersion"][0]
	// forceTranspile := form.Value["forceTranspile"][0]

	file, _ := c.FormFile("file")

	userWorkspacePath, err := utils.GetUserWorkspace(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := file.Filename
	extension := filepath.Base(fileName)
	userProj := path.Join(userWorkspacePath, "user-proj", "src")

	hasJsPrefix := strings.Contains(extension, ".js")

	if hasJsPrefix {
		c.SaveUploadedFile(file, userProj+"/views/"+fileName)
	} else {
		c.SaveUploadedFile(file, userProj+"/"+fileName)
	}

	// // TODO - Transpile JSX -> JS
	err = lib.PostFile(userProj, userWorkspacePath, filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload feito com sucesso",
	})

}
