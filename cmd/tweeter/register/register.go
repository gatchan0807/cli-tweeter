package register

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"log"
)

func Register(context *cli.Context) error {

	fmt.Println("register")

	if _, err := os.Stat("/tmp/tweeter/user_account.txt"); os.IsNotExist(err) {
		os.Mkdir("/tmp/tweeter", os.ModePerm)
		_, err := os.Create("/tmp/tweeter/user_account.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}
