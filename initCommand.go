package main

import (
	"github.com/urfave/cli"
	"os"
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

	createDirectoryIfNotExists(homeDir + "/.slurp")
	createDirectoryIfNotExists(homeDir + "/.slurp/projects")

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

func createDirectoryIfNotExists(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, os.ModePerm)
	}
}
