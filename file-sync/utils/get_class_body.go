package utils

import (
	"regexp"
	"strings"
)

func GetClassBody(jsContent string) string {
	pattern1 := regexp.MustCompile(`[^\{]+{`)
	pattern2 := regexp.MustCompile(`\}\s*$`)
	foundPtt1 := pattern1.FindString(jsContent)
	foundPtt2 := pattern2.FindString(jsContent)

	content := strings.Replace(jsContent, foundPtt1, "", 1)
	content = strings.Replace(content, foundPtt2, "", 1)

	return content
}
