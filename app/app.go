package app

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/imemo/domain/service/commands"
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
	app.Commands = commands.GetCommandMap()
	// command validator registration
	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	return app
}
