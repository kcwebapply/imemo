package app

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/imemo/commands"
	validator "github.com/kcwebapply/imemo/validator"
)

var appName = "imemo"
var version = "1.0.1"

// InitApp method is for initializing and  getting App settings.
func InitApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "github.com/kcwebapply/imemo"
	app.Version = version

	// dry run config
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d",
			Usage: "for dry run",
		},
	}
	// command routing.
	app.Commands = []cli.Command{
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
	// command validator registration
	app.Before = func(c *cli.Context) error {
		err := validator.ValidateArgument(c)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	return app
}
