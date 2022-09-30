package utils

import (
	"os"
	"path"
)

func GetReactProj() (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {

		return "", err
	}
	return path.Join(currentPath, "..", "react-proj/"), nil
}
