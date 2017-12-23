package register

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/checker"
	"io/ioutil"
	"strings"
)

func Register(context *cli.Context) error {

	fmt.Println("register")

	if _, err := os.Stat("/tmp/tweeter/user_account.csv"); os.IsNotExist(err) {
		os.Mkdir("/tmp/tweeter", os.ModePerm)
		_, err := os.Create("/tmp/tweeter/user_account.csv")
		checker.Check(err)
	}

	if isExist("ahaha") {
		fmt.Println(context.Args().Get(0) + " is already exist.")
	}

	return nil
}

func isExist(userId string) bool {
	data, err := ioutil.ReadFile("/tmp/tweeter/user_account.csv")
	checker.Check(err)

	userIdList := convertToUserIdList(string(data))

	for _, element := range userIdList {
		if element == userId {
			return true
		}
	}
	return false
}

func convertToUserIdList(rawData string) []string {
	var result []string

	perLine := strings.Split(rawData, "\n")

	for _, element := range perLine {
		result = append(result, strings.Split(element, ",")[0])
	}

	return result
}
