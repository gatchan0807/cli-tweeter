package list

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
)

func List(context *cli.Context) error {

	userInfoList := util.GetUserInfoList()

	for _, element := range userInfoList {
		fmt.Println(element["userId"])
	}

	return nil
}
