package utils

import "os"

func ReadFileAsString(filePath string) string {
	fileBuf, _ := os.ReadFile(filePath)

	return string(fileBuf)
}
