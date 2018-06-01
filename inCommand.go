package main

import "github.com/urfave/cli"

func inCommandAction(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}