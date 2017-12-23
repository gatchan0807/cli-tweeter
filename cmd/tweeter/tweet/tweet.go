package tweet

import (
	"github.com/urfave/cli"
	"fmt"
)

func Tweet(context *cli.Context) error {

	if context.String("account") != "" {
		TweetWithAccount(context, context.String("account"))
		return nil
	}

	fmt.Println("tweet")
	fmt.Println(context.Args().Get(0))

	return nil
}

func TweetWithAccount(context *cli.Context, accountId string) error {

	fmt.Println("tweet with account")
	fmt.Println(accountId)

	fmt.Println(context.Args().Get(0))

	return nil
}
