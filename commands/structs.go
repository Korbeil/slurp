package commands

// Project is used to store project details
type Project struct {
	Name      string
	Directory string
}

// Environment is used to store context data
type Environment struct {
	Project string
	OldBashHistoryPath string
}