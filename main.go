package main

import (
	"github.com/Josgtg/git-clone/cli"
)

// TODO: Create tests for the object types

func main() {
	args := cli.ParseArgs()

	switch args.Command() {
	case "init":
	case "add":
	case "status":
	case "commit":
	default:
		panic("unrecognized argument")
	}
}
