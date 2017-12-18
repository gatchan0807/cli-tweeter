package register

import (
	"github.com/urfave/cli"
	"fmt"
)

func Register(context *cli.Context) error {

	fmt.Println("register")
	fmt.Println(context.Args().Get(0))
	return nil
}
