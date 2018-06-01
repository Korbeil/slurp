package main

import "github.com/urfave/cli"

func makeInitCommand() cli.Command {
	return cli.Command{
		Name:      "init",
		Aliases:   []string{"i"},
		Usage:     "Initialize new project in current directory, if you give no project name, it will make a slug of directory name.",
		ArgsUsage: "<project>",
		Action: func(c *cli.Context) error {
			print("init")
			return nil
		},
	}
}
