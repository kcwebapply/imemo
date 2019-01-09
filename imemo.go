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
		// コマンド設定
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "メモ一覧を表示します。",
			Action:  commands.GetAllMemo,
		},
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "メモを保存します。",
			Action:  commands.SaveMemo,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "指定したidを持つメモを削除します。",
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
