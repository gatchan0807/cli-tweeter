package util

import (
	"strings"
)

func ConvertToUserIdList(rawData string) []string {
	var result []string

	perLine := strings.Split(rawData, "\n")
	perLine = perLine[:len(perLine)-1] // 末尾削除

	for _, element := range perLine {
		result = append(result, strings.Split(element, ",")[0])
	}

	return result
}
