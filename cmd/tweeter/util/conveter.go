package util

import (
	"strings"
	"io/ioutil"
)

var accountListFilePath = "/tmp/tweeter/user_account.csv"

func ConvertToIdList(rawData string) []string {
	var result []string

	perLine := strings.Split(rawData, "\n")
	perLine = perLine[:len(perLine)-1] // 末尾削除

	for _, element := range perLine {
		result = append(result, strings.Split(element, ",")[0])
	}

	return result
}

func GetUserInfoList() []map[string]string {
	var result []map[string]string

	rawData, err := ioutil.ReadFile(accountListFilePath)
	Check(err)

	perLine := strings.Split(string(rawData), "\n")
	perLine = perLine[:len(perLine)-1] // 末尾削除

	for _, element := range perLine {
		rawData := strings.Split(element, ",")
		rawMap := map[string]string{
			"userId":       rawData[0],
			"accessToken":  rawData[1],
			"accessSecret": rawData[2],
		}
		result = append(result, rawMap)
	}

	return result
}
