package directory

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func UserHome() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func Current() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return path.Base(pwd)
}

func CreateButWarnIfExists(directory string, message string) {
	if _, err := os.Stat(directory); err == nil {
		fmt.Printf(message, directory)
		os.Exit(1)
	}

	os.MkdirAll(directory, os.ModePerm)
}
