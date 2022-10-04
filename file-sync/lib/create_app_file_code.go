package lib

import (
	"fmt"
	"os"
	"path"
	"strings"
	"workspace_go/main/utils"
)

func CreateAppFileCode(basename string, files []string, folderPath string) {
	currentPath, _ := os.Getwd()
	templatePath := path.Join(currentPath, "template", "reactAppFile.txt")
	templateData := utils.ReadFileAsString(templatePath)

	imports := generateImports(files)
	templateData = strings.Replace(templateData, "#IMPORTS", imports, 1)
	templateData = strings.Replace(templateData, "#BROWSER_ROUTER", createBrowserRouter(files), 1)
	go utils.CreateFile(basename, templateData)
}

func generateImports(files []string) string {
	imports := ""

	for _, file := range files {
		if file != "" {
			importFile := fmt.Sprintf("import %s from './views/%s';\n", file, file)
			imports += importFile
		}
	}

	return imports
}

func createRoute(name string) string {
	if name != "" {
		return fmt.Sprintf("<Route path={['/%s', '/%s.html']} component={this.state.errorInfo ? ErrorInfoComponent : %s} />\n", name, name, name)
	}
	return ""
}

func createRoutes(names []string) string {
	routes := ""
	for _, name := range names {
		route := createRoute(name)
		routes += route
	}

	return routes
}

func createBrowserRouter(files []string) string {
	indexTag := ""
	index := indexOf("Index", files)

	if index != -1 {
		files = files[:index]
		indexTag = `<Route component={this.state.errorInfo ? ErrorInfoComponent : Index} />`
	}

	if indexTag == "" {
		index = indexOf("Home", files)
		if index != -1 {
			files = files[:index]
			indexTag = `<Route component={this.state.errorInfo ? ErrorInfoComponent : Home} />`
		}
	}
	routes := createRoutes(files)
	return fmt.Sprintf(`<Router history={history}>
	<Switch>
			%s
			%s
	</Switch>
</Router>`, routes, indexTag)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
