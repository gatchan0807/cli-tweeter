package list

import (
	"github.com/urfave/cli"
	"fmt"
	"io/ioutil"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
)

func List(context *cli.Context) error {

	fmt.Println("list")

	data, err := ioutil.ReadFile("/tmp/tweeter/user_account.csv")
	if err != nil {
		fmt.Println("ユーザーのアカウントは登録されていませんでした。")
		return nil
	}

	userIdList := util.ConvertToUserIdList(string(data))

	for _, element := range userIdList {
		fmt.Println(element)
	}

	return nil
}
