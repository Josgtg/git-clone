package cli

import "github.com/alecthomas/kong"

// Struct representing the arguments needed to be passed to the program.
// Intented to be used with the Kong parser.
var CLI struct {
	Init struct {} `cmd:""  help:"Initialize a repository."`

	Add struct {
		Files []string `arg:""  type:"path"  help:"Files to include on the next commit."`
	} `cmd:""  help:"Adds file(s) to be included in the next commit."`

	Status struct {} `cmd:""  help:"Shows status of the current commit."`

	Commit struct {
		Message string `short:"m"  help:"Message explaining the reason of the commit."  placeholder:"first commit"`
	} `cmd:""  help:"Realizes a commit including all the changes added."`
}

func GetArgs() *kong.Context {
	return kong.Parse(&CLI)
} 

