package commands

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/imemo/domain/repository"
	"github.com/kcwebapply/imemo/util"
)

func Clear(c *cli.Context) {

	if err := repository.DeleteAllMemo(); err != nil {
		util.HandleError(err)
	}

	if err := repository.DeleteAllCategory(); err != nil {
		util.HandleError(err)
	}
}
