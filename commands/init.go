package commands

import (
	"github.com/Korbeil/slurp/utils/directory"
	"github.com/Korbeil/slurp/utils/json"
	"github.com/urfave/cli"
	"path"
)

// InitCommand is used to create cli.Command object
// Also, this command is used to create a new project
func InitCommand() cli.Command {
	return cli.Command{
		Name:      "init",
		Aliases:   []string{"i"},
		Usage:     "Initialize new project in current utils, if you give no project name, it will make a slug of utils name.",
		ArgsUsage: "[project]",
		Action:    makeInitAction,
	}
}

func makeInitAction(c *cli.Context) error {
	homeDir := directory.UserHome()
	currentDir := directory.Current()

	projectName := c.Args().First()
	if projectName == "" {
		projectName = path.Base(currentDir)
	}

	directory.CreateButWarnIfExists(
		homeDir+"/.slurp/projects/"+projectName,
		"Project utils with path `%s` already exists.\n")

	json.WriteJsonInFile(
		Project{Name: projectName, Directory: currentDir},
		homeDir+"/.slurp/projects/"+projectName+"/config.json")

	return nil
}
