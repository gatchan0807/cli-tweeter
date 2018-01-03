package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/user"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/tweet"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/list"
)

func main() {
	app := cli.NewApp()

	app.Name = "cli tweeter"
	app.Usage = "first, register twitter account. second, let's tweet!"
	app.Version = "0.0.1"

	app.Action = tweet.Tweet

	app.Commands = []cli.Command{
		{
			Name:    "user",
			Aliases: []string{"users"},
			Usage:   "Register twitter account information",
			Action:  user.User,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "delete, d",
				},
			},
		},
		{
			Name:    "tweet",
			Aliases: []string{"t"},
			Usage:   "Do tweet",
			Action:  tweet.Tweet,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "account, a",
				},
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List up registered accounts",
			Action:  list.List,
		},
	}

	app.Run(os.Args)
}
