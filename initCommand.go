package main

import (
	"github.com/Korbeil/slurp/utils/directory"
	"github.com/urfave/cli"
)

func makeInitCommand() cli.Command {
	return cli.Command{
		Name:      "init",
		Aliases:   []string{"i"},
		Usage:     "Initialize new project in current utils, if you give no project name, it will make a slug of utils name.",
		ArgsUsage: "<project>",
		Action:    makeInitAction,
	}
}

func makeInitAction(c *cli.Context) error {
	homeDir := directory.UserHome()
	projectDirName := directory.Current()

	projectName := c.Args().First()
	if projectName == "" {
		projectName = projectDirName
	}

	directory.CreateButWarnIfExists(
		homeDir+"/.slurp/projects/"+projectName,
		"Project utils with path `%s` already exists.\n")

	return nil
}
