package lib

import (
	"os"
	"path"
	"strings"
	"workspace_go/main/utils"
)

func Template(className string, jsContent string, htmlContent string, ameComponents string) string {
	currentPath, _ := os.Getwd()
	templatePath := path.Join(currentPath, "template", "reactFile.txt")
	templateData := utils.ReadFileAsString(templatePath)

	templateData = strings.Replace(templateData, "#TAGS", ameComponents, 1)
	templateData = strings.Replace(templateData, "#CLASS_NAME", className, 1)
	templateData = strings.Replace(templateData, "#JSX_CONTENT", htmlContent, 1)
	templateData = strings.Replace(templateData, "#JS_CONTENT", jsContent, 1)

	return templateData
}
