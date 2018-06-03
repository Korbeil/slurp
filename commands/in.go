package commands

import (
	"github.com/urfave/cli"
	"github.com/Korbeil/slurp/utils/directory"
	"os"
	"fmt"
)

// InCommandAction is used as default action
// Also, it permits to jump from one project to another
func InCommandAction(c *cli.Context) error {
	homeDir := directory.UserHome()
	projectName := c.Args().First()
	env := loadEnv(homeDir)

	// Checking if env is empty 
	if env.Project == "" {
		fmt.Printf("No last project in env.")
		os.Exit(1)
	}
	// If no project name given, we take last project we loaded from env
	if projectName == "" {
		projectName = env.Project
	}

	project := loadProject(homeDir, projectName)

	// Setting new bash_history
	os.Setenv("HISTFILE", homeDir+"/.slurp/projects/"+projectName+"/bash_history")

	print(project.Directory)

	return nil
}
