package util

import (
	"log"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FindUserId(userId string) map[string]string {
	userInfoList := GetUserInfoList()

	for _, element := range userInfoList {
		if element["userId"] == userId {
			return element
		}
	}
	return nil
}
