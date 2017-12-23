package register

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/checker"
)

func Register(context *cli.Context) error {

	fmt.Println("register")

	if _, err := os.Stat("/tmp/tweeter/user_account.csv"); os.IsNotExist(err) {
		os.Mkdir("/tmp/tweeter", os.ModePerm)
		_, err := os.Create("/tmp/tweeter/user_account.csv")
		checker.Check(err)
	}

	return nil
}
