package commands

import (
	"errors"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/kcwebapply/imemo/domain/repository"
	"github.com/kcwebapply/imemo/util"
	"github.com/kcwebapply/imemo/view"
)

// DeleteMemo delete memodata
func Rm(c *cli.Context) {
	memoID, err := strconv.Atoi(c.Args().First())
	if err != nil {
		util.HandleError(errors.New("argument error. first paramerter should be int-format"))
	}

	// get memo object begore hand to get categoryId
	memo, err := repository.GetMemo(memoID)
	categoryID := memo.Categoryid
	if err != nil {
		util.HandleError(err)
	}

	if err := repository.DeleteMemoById(memoID); err != nil {
		util.HandleError(err)
	}

	// delete category if no memo are belonged to.
	memos, err := repository.GetMemosByCategoryId(categoryID)
	if err != nil {
		util.HandleError(err)
	}

	if len(memos) == 0 {
		// delete category.
		repository.DeleteCategoryById(categoryID)
	}

	view.PrintDeleteMemo(memo)

}
