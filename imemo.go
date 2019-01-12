package main

import (
	"os"

	"github.com/codegangsta/cli"
	app "github.com/kcwebapply/imemo/app"
	commands "github.com/kcwebapply/imemo/commands"
)

func main() {
	app := app.GetApp()
	app.Commands = []cli.Command{
		// Configuration for command
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "View saved memo.",
			Action:  commands.GetAllMemo,
		},
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "Save memo.",
			Action:  commands.SaveMemo,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete memo that specified by id.",
			Action:  commands.DeleteMemo,
		},

		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "Edit memo that specified by id.",
			Action:  commands.EditMemo,
		},

		{
			Name:    "clear",
			Aliases: []string{"c"},
			Usage:   "clear (delete) all memo data.",
			Action:  commands.ClearMemo,
		},
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}

	app.Run(os.Args)
}
