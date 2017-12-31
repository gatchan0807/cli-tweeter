package list

import (
	"github.com/urfave/cli"
	"fmt"
	"io/ioutil"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
)

var accountListFilePath = "/tmp/tweeter/user_account.csv"

func List(context *cli.Context) error {

	data, err := ioutil.ReadFile(accountListFilePath)
	if err != nil {
		fmt.Println("ユーザーのアカウントは登録されていませんでした。")
		fmt.Println("(User account not found.)")
		return nil
	}

	userIdList := util.ConvertToIdList(string(data))

	for _, element := range userIdList {
		fmt.Println(element)
	}

	return nil
}
