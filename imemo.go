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
	}

	app.Before = func(c *cli.Context) error {
		// 開始前の処理をここに書く
		//fmt.Println("開始")
		return nil // error を返すと処理全体が終了
	}

	app.After = func(c *cli.Context) error {
		// 終了時の処理をここに書く
		//fmt.Println("終了")
		return nil
	}

	app.Run(os.Args)
}
