package utils

import (
	"path/filepath"
	"strings"
)

func Is2Process(relativePath string) bool {
	ext := filepath.Ext(relativePath)
	isViews := strings.Contains(relativePath, "/views/")
	isComponents := strings.Contains(relativePath, "/components/")

	if (ext == ".js" || ext == ".html" || ext == ".jsx") && (isViews || isComponents) {
		return true
	}
	return false
}
