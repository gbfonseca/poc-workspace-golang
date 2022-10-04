package lib

import (
	"os"
	"path"
	"workspace_go/main/utils"
)

func PostFile(userProjPath string, userWorkspacePath string, filePath string) error {
	userReactProj := path.Join(userWorkspacePath, "react-proj")

	absolutePath := path.Join(userProjPath, filePath)
	fileBuf, _ := os.ReadFile(absolutePath)

	data := string(fileBuf)

	if utils.Is2Process(filePath) {
		err := CreateReactFile(userProjPath, userReactProj, data, filePath, absolutePath)
		if err != nil {
			return err
		}
	}

	return nil

}
