# slurp

[![codebeat badge](https://codebeat.co/badges/465043f5-f438-4fe8-a99a-c1312e76ebe4)](https://codebeat.co/projects/github-com-korbeil-slurp-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Korbeil/slurp)](https://goreportcard.com/report/github.com/Korbeil/slurp)

## Get started

To have it usable through your `$GOPATH` (check [GOPATH official wiki](https://github.com/golang/go/wiki/GOPATH) page for more details), you simply have to run theses commands:
```bash
make deps
make build
make install
```

Like that, you'll have a working `slurp` command within your terminal :)

## Usage

### slurp project

It permits to jump from one project to another.
The project argument is required and is also the name of the project to jump to.

### slurp init [project]

This command is used to create a new project.
The project argument is optional and is also the name of the project to create, if not given we take the current directory slugged name.

### slurp burp

This command is used to reset all slurp behaviors.
