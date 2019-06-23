package commands

import (
	"github.com/codegangsta/cli"
)

var fileName = ""

var maxTextSize = 60

func GetCommandMap() []cli.Command {
	var commands = []cli.Command{
		{
			Name:    "ls",
			Aliases: []string{},
			Usage:   "list memo. ",
			Action:  Ls,
		},
		{
			Name:    "add",
			Aliases: []string{"add"},
			Usage:   "Save memo.",
			Action:  Add,
		},
		{
			Name:    "rm",
			Aliases: []string{"r"},
			Usage:   "Delete memo that specified by id.",
			Action:  Rm,
		},
		{
			Name:    "clear",
			Aliases: []string{"c"},
			Usage:   "clear (delete) all memo data.",
			Action:  Clear,
		},
	}
	return commands
}
