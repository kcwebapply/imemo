package app

import (
	"github.com/codegangsta/cli"
	//. "github.com/kcwebapply/imemo/config"
)

var appName string
var version string

func init() {
	//config := GetConfig()
	//appName = config.App.Name
	//version = config.App.Version
	appName = "imemo"
	version = "1.0.0"
}

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "github.com/kcwebapply/imemo"
	app.Version = version

	// グローバルオプション設定
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dryrun, d",
			Usage: "for dry run",
		},
	}
	return app
}
