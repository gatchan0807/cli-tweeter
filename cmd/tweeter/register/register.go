package register

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"io/ioutil"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
)

func Register(context *cli.Context) error {

	fmt.Println("register")

	if _, err := os.Stat("/tmp/tweeter/user_account.csv"); os.IsNotExist(err) {
		os.Mkdir("/tmp/tweeter", os.ModePerm)
		_, err := os.Create("/tmp/tweeter/user_account.csv")
		util.Check(err)
	}

	if isExist(context.Args().Get(0)) {
		fmt.Println(context.Args().Get(0) + " is already exist.")
	}

	//

	return nil
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
