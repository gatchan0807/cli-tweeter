package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/ahaha0807/cli-tweeter/cmd/tweeter/register"
)

func main() {
	app := cli.NewApp()

	app.Name = "cli tweeter"
	app.Usage = "first, register twitter account. second, let's tweet!"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) {
		cli.ShowAppHelp(context)
	}

	app.Commands = []cli.Command{
		{
			Name:    "register",
			Aliases: []string{"r"},
			Usage:   "Register twitter account info.",
			Action:  register.Register,
		},
	}

	app.Run(os.Args)
}
