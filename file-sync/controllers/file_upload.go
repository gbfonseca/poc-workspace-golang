package controllers

import (
	"net/http"
	"path"
	"workspace_go/main/utils"

	"github.com/gin-gonic/gin"
)

func FileUpload(c *gin.Context) {
	form, _ := c.MultipartForm()

	email := form.Value["email"][0]
	// filePath := form.Value["filepath"][0]
	// why := form.Value["why"][0]
	// miniAppVersion := form.Value["miniAppVersion"][0]
	// forceTranspile := form.Value["forceTranspile"][0]

	file, _ := c.FormFile("file")

	userWorkspacePath, err := utils.GetUserWorkspace(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// extension := filepath.Base()(file.Filename)
	// newFileName := "miniapp" + extension

	userProj := path.Join(userWorkspacePath, "user-proj", "src")

	c.SaveUploadedFile(file, userProj+"/"+file.Filename)

	// err = utils.Unzip(userProj+"/"+newFileName, userProj)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// cmd := exec.Command("rm", "-rf", userProj+"/"+newFileName)
	// go cmd.Run()

	// // TODO - Transpile JSX -> JS
	// err = lib.PostFiles(userProj, userWorkspacePath)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload feito com sucesso",
	})

}
