package main

import "github.com/urfave/cli"

func makeConsoleApplication() *cli.App {
	app := cli.NewApp()
	app.Name = "slurp"
	app.Usage = "Simple multiple project manager"
	app.Version = "1.0.0-alpha"
	app.Copyright = "(c) 2018 Baptiste Leduc"
	app.UsageText = "slurp <project> \n\t slurp init <project>\n\t slurp burp"

	app.Commands = []cli.Command{
		{
			Name:      "init",
			Aliases:   []string{"i"},
			Usage:     "Initialize new project in current directory, if you give no project name, it will make a slug of directory name.",
			ArgsUsage: "<project>",
			Action: func(c *cli.Context) error {
				print("init")
				return nil
			},
		},
		{
			Name:    "burp",
			Aliases: []string{"b"},
			Usage:   "Leaving project: reset slurp aliases and bash history paths.",
			Action: func(c *cli.Context) error {
				print("burp")
				return nil
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "verbose, V",
			Usage: "make the binary more verbose",
		},
	}

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	return app
}