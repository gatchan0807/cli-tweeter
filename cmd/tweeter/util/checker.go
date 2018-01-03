package util

import (
	"log"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FindUserInfo(userId string) map[string]string {
	userInfoList := GetUserInfoList()

	for _, element := range userInfoList {
		if element["userId"] == userId {
			return element
		}
	}
	return nil
}

func FindUserIndex(userId string) int {
	userInfoList := GetUserInfoList()

	for index, element := range userInfoList {
		if element["userId"] == userId {
			return index
		}
	}

	return -1
}