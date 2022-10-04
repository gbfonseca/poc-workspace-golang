package lib

import (
	"bufio"
	"os"
	"path"
	"strings"
	"workspace_go/main/utils"
)

func PostFile(userProjPath string, userWorkspacePath string) error {
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

	filesReadOrder = utils.RemoveDuplicateStr(filesReadOrder)
	for _, fileRelativePath := range filesReadOrder {
		absolutePath := path.Join(userProjPath, fileRelativePath)
		fileBuf, _ := os.ReadFile(absolutePath)

		data := string(fileBuf)

		if utils.Is2Process(fileRelativePath) {
			err := CreateReactFile(userProjPath, userReactProj, data, fileRelativePath, absolutePath)
			if err != nil {
				return err
			}
		}

	}

	return nil

}
