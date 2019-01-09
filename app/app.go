package app

import (
	"github.com/codegangsta/cli"
)

var appName = "imemo"
var version = "1.2.0"

// GetApp method is for getting App settings.
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
