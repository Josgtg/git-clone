package main

import (
	"github.com/Josgtg/git-clone/cli"
)

func main() {
	args := cli.GetArgs()

	switch args.Command() {
	case "init":
	case "add":
	case "status":
	case "commit":
	default:
		panic("Unrecognized argument")
	}
}
