package commands

import "github.com/urfave/cli"

func InCommandAction(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}
