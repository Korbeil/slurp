package commands

import (
	"os"
	"github.com/urfave/cli"
	"github.com/Korbeil/slurp/utils/directory"
)

// BurpCommand is used to create cli.Command object
// Also, this command is used to reset all slurp behaviors
func BurpCommand() cli.Command {
	return cli.Command{
		Name:    "burp",
		Aliases: []string{"b"},
		Usage:   "Leaving project: reset slurp aliases and bash history paths.",
		Action:  makeOutAction,
	}
}

func makeOutAction(c *cli.Context) error {
	homeDir := directory.UserHome()
	env := loadEnv(homeDir)

	// Setting new bash_history
	os.Setenv("HISTFILE", env.OldBashHistoryPath)

	return nil
}
