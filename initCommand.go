package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"path"
	"runtime"
)

func makeInitCommand() cli.Command {
	return cli.Command{
		Name:      "init",
		Aliases:   []string{"i"},
		Usage:     "Initialize new project in current directory, if you give no project name, it will make a slug of directory name.",
		ArgsUsage: "<project>",
		Action:    makeInitAction,
	}
}

func makeInitAction(c *cli.Context) error {
	homeDir := userHomeDir()
	projectDirName := currentProjectDir()

	projectName := c.Args().First()
	if projectName == "" {
		projectName = projectDirName
	}

	projectDir := homeDir + "/.slurp/projects/" + projectName
	createDirectoryWarnIfExists(projectDir)

	return nil
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func currentProjectDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path.Base(pwd)
}

func createDirectoryWarnIfExists(directory string) {
	if _, err := os.Stat(directory); err == nil {
		fmt.Printf("Project directory with path `%s` already exists.\n", directory)
		os.Exit(1)
	}

	os.MkdirAll(directory, os.ModePerm)
}
