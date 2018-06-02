package directory

import (
	"fmt"
	"os"
	"runtime"
)

// Get user home fullpath
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

// Get current directory fullpath
func Current() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

// Create directory or exit with message if already exists
func CreateButWarnIfExists(directory string, message string) {
	if _, err := os.Stat(directory); err == nil {
		fmt.Printf(message, directory)
		os.Exit(1)
	}

	os.MkdirAll(directory, os.ModePerm)
}
