package list

import (
	"fmt"

	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/util"
	"github.com/urfave/cli"
)

// Display registered user ids.
func List(_ *cli.Context) error {

	userInfoList := util.GetUserInfoList()

	for _, element := range userInfoList {
		fmt.Println(element["userId"])
	}

	return nil
}
