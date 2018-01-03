package list

import (
	"fmt"

	"github.com/ahaha0807/cli-tweeter/util"
	"github.com/urfave/cli"
)

// Display registered account ids.
func List(_ *cli.Context) error {

	userInfoList := util.GetUserInfoList()

	for _, element := range userInfoList {
		fmt.Println(element["userId"])
	}

	return nil
}
