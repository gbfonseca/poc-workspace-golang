package utils

import (
	"os"
	"path"
)

func GetUserWorkspace(email string) (string, error) {
	currentPath, err := os.Getwd()

	if err != nil {
		return "", err
	}

	return path.Join(currentPath, "..", "storage", email), nil
}
