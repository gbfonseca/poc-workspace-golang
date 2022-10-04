package utils

import (
	"regexp"
	"strings"
)

func TagsFinder(htmlContent string) string {

	m := regexp.MustCompile(`<(\w+)[\s(/>)]{1}`)

	tagsUnformatted := m.FindAllString(htmlContent, -1)
	var tags string

	for _, tag := range tagsUnformatted {
		tag = strings.Replace(tag, "<", "", 1)
		tag = strings.Replace(tag, ">", ",", 1)
		tags += tag
	}

	return tags

}
