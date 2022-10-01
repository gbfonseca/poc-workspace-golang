package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
	"workspace_go/main/models"
	"workspace_go/main/utils"

	"github.com/samber/lo"

	"github.com/gin-gonic/gin"
)

func Setup(c *gin.Context) {

	var ameConf models.AmeConf

	if err := c.ShouldBindJSON(&ameConf); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentPath, err := utils.GetReactProj()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userPath, _ := utils.GetUserWorkspace(ameConf.Email)

	exists := projectExists(userPath)

	if !exists {
		err := createSources(userPath, currentPath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	err = appendConfToPackageJson(ameConf, userPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

func appendConfToPackageJson(ameConf models.AmeConf, userProjPath string) error {
	pkgJsonPath := path.Join(userProjPath, "react-proj", "package.json")
	file, err := os.ReadFile(pkgJsonPath)
	if err != nil {
		return err
	}

	var pkgJson models.PackageJson
	err = json.Unmarshal(file, &pkgJson)
	if err != nil {
		return err
	}

	pkgJson.Name = ameConf.Slug
	pkgJson.Version = ameConf.Version

	dependencies := make(map[string]string)
	dependencies["ame-miniapp-components"] = ameConf.AmeMiniappComponents
	dependencies["ame-super-app-client"] = ameConf.AmeSuperAppClient

	updatedDependencies := lo.Assign(pkgJson.Dependencies, dependencies)

	pkgJson.Dependencies = updatedDependencies

	pkgData, _ := json.MarshalIndent(pkgJson, "", " ")

	err = os.WriteFile(pkgJsonPath, pkgData, os.ModeSticky|os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
