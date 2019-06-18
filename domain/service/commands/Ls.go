package commands

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/imemo/domain/model"
	"github.com/kcwebapply/imemo/domain/repository"
	"github.com/kcwebapply/imemo/view"
)

func Ls(c *cli.Context) {
	var memolist []model.Memo

	memolist, err := repository.GetMemos()

	if err != nil {
		fmt.Printf("db error! %s\n", err.Error())
	}

	view.PrintAllMemo(memolist)
}
