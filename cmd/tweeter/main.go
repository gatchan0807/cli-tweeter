package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "cli tweeter"
	app.Usage = "first, register twitter account. second, let's tweet!"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) {
		cli.ShowAppHelp(context)
	}

	app.Run(os.Args)
}
