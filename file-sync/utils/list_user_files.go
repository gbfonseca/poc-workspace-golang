package utils

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

func ListUserFiles(userProjPath string, directory string) ([]string, error) {
	directoryPath := path.Join(userProjPath, directory)

	filesInfo, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return nil, nil
	}

	var files []string

	for _, file := range filesInfo {
		fileName := file.Name()
		fileExt := filepath.Ext(fileName)

		fileName = strings.Replace(fileName, fileExt, "", 1)

		files = append(files, fileName)
	}

	files = RemoveDuplicateStr(files)

	return files, nil
}
