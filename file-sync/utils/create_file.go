package utils

import "os"

func CreateFile(filePath string, fileContent string) error {
	fileByte := []byte(fileContent)
	err := os.WriteFile(filePath, fileByte, os.ModeSticky|os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
