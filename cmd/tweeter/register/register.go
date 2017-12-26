package register

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
)

func Register(context *cli.Context) error {
	if _, err := os.Stat("/tmp/tweeter/user_account.csv"); os.IsNotExist(err) {
		os.Mkdir("/tmp/tweeter", os.ModePerm)
		_, err := os.Create("/tmp/tweeter/user_account.csv")
		util.Check(err)
	}

	userName := getUserId()
	userAccountToken := getTwitterToken()

	success := addToCsvFile(userName, userAccountToken)
	if !success {
		fmt.Println("Save failed")
	}

	return nil
}

func getUserId() string {
	var userAccountName string
	fmt.Println("Input your Twitter account ID.(without '@')")
	fmt.Scan(&userAccountName)

	for userAccountName == "" || isExist(userAccountName) {
		if isExist(userAccountName) {
			fmt.Println(userAccountName + " is already exist.")
		}

		fmt.Println("Input your Twitter account ID.(without '@')")
		fmt.Scan(&userAccountName)
	}

	return userAccountName
}

func isExist(userId string) bool {
	data, err := ioutil.ReadFile("/tmp/tweeter/user_account.csv")
	util.Check(err)

	userIdList := util.ConvertToUserIdList(string(data))

	for _, element := range userIdList {
		if element == userId {
			return true
		}
	}
	return false
}

func addToCsvFile(accountName string, accountToken string) bool {
	file, err := os.OpenFile("/tmp/tweeter/user_account.csv", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return false
	}

	writeLine := accountName + "," + accountToken + "\n"

	fmt.Fprint(file, writeLine)

	return true
}

func getTwitterToken() string {
	token := ""

	return token
}
