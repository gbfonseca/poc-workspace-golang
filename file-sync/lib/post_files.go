package lib

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func PostFiles(userProjPath string, userWorkspacePath string) error {
	userReactProj := path.Join(userWorkspacePath, "react-proj")
	filesOrderPath := path.Join(userProjPath, "__files_order.txt")

	buf, err := os.Open(filesOrderPath)
	if err != nil {
		return err
	}

	fileScan := bufio.NewScanner(buf)
	var filesReadOrder []string

	for fileScan.Scan() {
		fileText := strings.Replace(fileScan.Text(), ".jsx", ".js", 1)
		filesReadOrder = append(filesReadOrder, fileText)
	}

	filesReadOrder = removeDuplicateStr(filesReadOrder)
	for _, fileRelativePath := range filesReadOrder {
		absolutePath := path.Join(userProjPath, fileRelativePath)
		fileBuf, _ := os.ReadFile(absolutePath)

		data := string(fileBuf)

		if is2Process(fileRelativePath) {
			err := CreateReactFile(userProjPath, userReactProj, data, fileRelativePath, absolutePath)
			if err != nil {
				return err
			}
		}

	}

	return nil

}

func is2Process(relativePath string) bool {
	ext := filepath.Ext(relativePath)
	isViews := strings.Contains(relativePath, "/views/")
	isComponents := strings.Contains(relativePath, "/components/")

	if (ext == ".js" || ext == ".html" || ext == ".jsx") && (isViews || isComponents) {
		return true
	}
	return false
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
