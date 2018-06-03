package commands

import (
	"github.com/urfave/cli"
	utilsJson "github.com/Korbeil/slurp/utils/json"
	"github.com/Korbeil/slurp/utils/directory"
	"encoding/json"
)

// InCommandAction is used as default action
// Also, it permits to jump from one project to another
func InCommandAction(c *cli.Context) error {
	homeDir := directory.UserHome()
	projectName := c.Args().First()
	env := loadEnv(homeDir)

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
	var b []byte
	b = utilsJson.ReadJson(homeDir+"/.slurp/env.json")

	var env Environment
	json.Unmarshal(b, &env)
	return env
}
