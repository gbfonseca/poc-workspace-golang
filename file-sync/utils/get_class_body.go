package utils

import (
	"regexp"
	"strings"
)

func GetClassBody(jsContent string) string {
	pattern1 := regexp.MustCompile(`[^\{]+{`)
	foundPtt1 := pattern1.FindString(jsContent)
	content := strings.Replace(jsContent, foundPtt1, "", 1)
	content = content[:len(content)-3]

	return content
}
