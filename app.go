package main

import (
	"github.com/Korbeil/slurp/commands"
	"github.com/urfave/cli"
)

func makeConsoleApplication() *cli.App {
	app := cli.NewApp()
	app.Name = "slurp"
	app.Usage = "Simple multiple project manager"
	app.Version = "1.0.0-alpha"
	app.Copyright = "(c) 2018 Baptiste Leduc"
	app.UsageText = "slurp [project] \n\t slurp init <project>\n\t slurp burp"

	app.Commands = []cli.Command{
		commands.InitCommand(),
		commands.BurpCommand(),
	}

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "verbose, V",
			Usage: "make the binary more verbose",
		},
	}

	app.Action = commands.InCommandAction

	return app
}
