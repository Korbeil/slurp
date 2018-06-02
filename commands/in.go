package commands

import "github.com/urfave/cli"

// InCommandAction is used as default action
// Also, it permits to jump from one project to another
func InCommandAction(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return nil
}
