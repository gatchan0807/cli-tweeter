package tweet

import (
	"github.com/urfave/cli"
	"fmt"
)

func Tweet (context *cli.Context) error {

	fmt.Println("tweet")
	fmt.Println(context.Args().Get(0))

	return nil
}
