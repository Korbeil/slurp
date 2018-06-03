package commands

import (
	"github.com/urfave/cli"
	utilsJson "github.com/Korbeil/slurp/utils/json"
	"github.com/Korbeil/slurp/utils/directory"
	"encoding/json"
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

	print(project.Directory)

	return nil
}

func loadProject(homeDir string, projectName string) Config {
	var b []byte
	b = utilsJson.ReadJson(homeDir+"/.slurp/projects/"+projectName+"/config.json")

	var conf Config
	json.Unmarshal(b, &conf)
	return conf
}

func loadEnv(homeDir string) Environment {
	envPath := homeDir+"/.slurp/env.json"

	// env doesn't exists, let's create it and return an empty one
	if directory.CheckExists(envPath) == false {
		env := Environment{Project: "", OldBashHistoryPath: os.Getenv("HISTFILE")}
		utilsJson.WriteJsonInFile(
			env,
			envPath)
		return env
	}

	// or just read it :)
	var b []byte
	b = utilsJson.ReadJson(envPath)

	var env Environment
	json.Unmarshal(b, &env)
	return env
}
