package main

import (
	"os"

	app "github.com/kcwebapply/imemo/app"
)

func main() {
	app := app.InitApp()
	app.Run(os.Args)
}
