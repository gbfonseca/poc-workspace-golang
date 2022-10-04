package lib

import (
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"workspace_go/main/utils"
)

func CreateReactFile(userProjPath string, userReactProj string, data string, fileRelativePath string, absolutePath string) error {
	fileName := filepath.Base(absolutePath)
	fileExt := filepath.Ext(fileName)

	m := regexp.MustCompile(`\.js$`)

	userHtmlFileName := m.ReplaceAllString(fileName, ".jsx")
	userJsFileName := m.ReplaceAllString(fileName, ".js")
	userHtmlFilePath := strings.Replace(absolutePath, fileName, userHtmlFileName, 1)
	userJsFilePath := strings.Replace(absolutePath, fileName, userJsFileName, 1)

	className := strings.Replace(fileName, fileExt, "", 1)

	userHtmlContentData := utils.ReadFileAsString(userHtmlFilePath)
	userJsContentData := utils.ReadFileAsString(userJsFilePath)

	splitedJsContent := strings.Split(userJsContentData, "export default class ")

	userJsContentData = utils.GetClassBody(splitedJsContent[1])

	tags := utils.TagsFinder(userHtmlContentData)

	// files, err := utils.ListUserFiles(userProjPath, "src/components")
	// if err != nil {
	// 	return err
	// }

	template := Template(className, userJsContentData, userHtmlContentData, tags)

	reactFilePath := path.Join(userReactProj, "src", fileRelativePath)
	err := utils.CreateFile(reactFilePath, template)
	if err != nil {
		return err
	}
	appFilePath := path.Join(userReactProj, "src/App.js")

	err = CreateAppFile(appFilePath, userReactProj)
	if err != nil {
		return err
	}

	cmd := exec.Command("yarn", "esbuild")
	cmd.Dir = userReactProj

	cmd.Run()

	createHTMLFiles(userReactProj)

	return nil
}

func createHTMLFiles(userReactProj string) {
	srcEsbuildPath := path.Join(userReactProj, "src", "esbuild.html")
	buildIndexPath := path.Join(userReactProj, "build", "index.html")
	utils.Copy(srcEsbuildPath, buildIndexPath)
}
