package util

import (
	"strings"
	"io/ioutil"
	"log"
)

func GetUserInfoList() []map[string]string {
	var accountListFilePath = "/tmp/tweeter/user_account.csv"

	var result []map[string]string

	rawData, err := ioutil.ReadFile(accountListFilePath)
	if err != nil {
		log.Fatalf("ユーザー情報が見つかりませんでした。")
		log.Fatalf("(User account was not found.)")
		return nil
	}

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
