package commands

import "github.com/urfave/cli"

// BurpCommand is used to create cli.Command object
// Also, this command is used to reset all slurp behaviors
func BurpCommand() cli.Command {
	return cli.Command{
		Name:    "burp",
		Aliases: []string{"b"},
		Usage:   "Leaving project: reset slurp aliases and bash history paths.",
		Action: func(c *cli.Context) error {
			print("burp")
			return nil
		},
	}
}
