package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/register"
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
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "Register twitter account info.",
			Action:  register.Register,
		},
		{
			Name:    "tweet",
			Aliases: []string{"t"},
			Usage:   "Tweet a contents",
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
			Usage:   "List accounts information",
			Action:  list.List,
		},
	}

	app.Run(os.Args)
}
