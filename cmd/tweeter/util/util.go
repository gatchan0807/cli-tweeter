package util

import (
	"log"
	"io/ioutil"
	"strings"
)

// Error check method.
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Find user information from user list.
func FindUserInfo(userId string) map[string]string {
	userInfoList := GetUserInfoList()

	for _, element := range userInfoList {
		if element["userId"] == userId {
			return element
		}
	}
	return nil
}

// Find user index in user list.
func FindUserIndex(userId string) int {
	userInfoList := GetUserInfoList()

	for index, element := range userInfoList {
		if element["userId"] == userId {
			return index
		}
	}

	return -1
}

// Get user information list from csv file.
func GetUserInfoList() []map[string]string {
	var accountListFilePath = "/tmp/tweeter/accounts.csv"

	var result []map[string]string

	rawData, err := ioutil.ReadFile(accountListFilePath)
	if err != nil {
		log.Fatalf("アカウント情報が見つかりませんでした。")
		log.Fatalf("(Account was not found.)")
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
