package lib

import (
	"workspace_go/main/utils"
)

func CreateAppFile(basename string, folderPath string) error {
	files, err := utils.ListUserFiles(folderPath, "src/views")
	if err != nil {
		return err
	}

	CreateAppFileCode(basename, files, folderPath)

	return nil
}
