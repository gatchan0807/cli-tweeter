package util

import (
	"os"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

var accountListFilePath = "/tmp/tweeter/user_account.csv"

// Call this method then arguments data push to csv file.
func Push(accountName, accountToken, accountSecret string) error {
	file, err := os.OpenFile(accountListFilePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("File open Failed.")
	}

	writeLine := accountName + "," + accountToken + "," + accountSecret + "\n"
	fmt.Fprint(file, writeLine)

	return nil
}

// Replace data in csv file by arguments map.
func Replace(informationList []map[string]string) error {
	text := ""
	for _, element := range informationList {
		oneLine := element["userId"] + "," + element["accessToken"] + "," + element["accessSecret"]
		text += oneLine + "\n"
	}

	err := ioutil.WriteFile(accountListFilePath, []byte(text), 0666)
	if err != nil {
		return errors.New("File open Failed.")
	}

	return nil
}
