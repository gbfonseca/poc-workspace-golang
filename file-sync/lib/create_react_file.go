package lib

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"workspace_go/main/utils"
)

func CreateReactFile(userProjPath string, userReactProj string, data string, fileRelativePath string, absolutePath string) error {
	fileName := filepath.Base(absolutePath)
	fileExt := filepath.Ext(fileName)

	fileNameWithoutExt := strings.Replace(fileName, fileExt, "", 1)

	userHtmlFileName := fileNameWithoutExt + ".jsx"
	userJsFileName := fileNameWithoutExt + ".js"
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

	if fileExt == ".jsx" {
		fileRelativePath = strings.Replace(fileRelativePath, ".jsx", ".js", 1)
	}

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

	buildPath := path.Join(userReactProj, "build")

	files, _ := os.ReadDir(buildPath)

	if files != nil {
		cmd := exec.Command("rm", "-rf", "build")
		cmd.Dir = userReactProj
		cmd.Run()
	}

	cmd := exec.Command("yarn", "esbuild")
	cmd.Dir = userReactProj
	cmd.Run()

	createHTMLFiles(userReactProj)

	return nil
}

func createHTMLFiles(userReactProj string) {
	srcEsbuildPath := path.Join(userReactProj, "src", "esbuild.html")
	srcIndexPath := path.Join(userReactProj, "src", "index.html")
	buildIndexPath := path.Join(userReactProj, "build", "index.html")
	publicIndexPath := path.Join(userReactProj, "public", "index.html")

	utils.Copy(srcEsbuildPath, buildIndexPath)
	utils.Copy(srcIndexPath, publicIndexPath)

}
