package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/kcwebapply/imemo/domain/model"
	"github.com/kcwebapply/imemo/domain/repository"
	"github.com/kcwebapply/imemo/view"
)

func Add(c *cli.Context) {

	memostring := c.Args().Get(0)
	category_name := c.Args().Get(1)

	if memostring == "" {

	}

	if category_name == "" {
		add(memostring)
	} else {
		addWithCategory(memostring, category_name)
	}

}

func add(memostring string) {

	memo := model.Memo{Memo: memostring, Categoryid: 1}

	if err := repository.AddMemo(memo); err != nil {
		fmt.Println("memo error!")
	}

	view.PrintAddMemo(memo)
}

func addWithCategory(memostring string, category_string string) {

	categoryEntity := model.Category{Name: category_string}

	if err := repository.AddCategory(categoryEntity); err != nil {
		fmt.Println("category add error!")
		os.Exit(0)
	}

	categoryEntity, err := repository.GetCategoryByName(category_string)
	if err != nil {
		fmt.Println("category get error!")
		os.Exit(0)
	}

	id := categoryEntity.ID

	memo := model.Memo{Memo: memostring, Categoryid: id}

	if err := repository.AddMemo(memo); err != nil {
		fmt.Println("memo error!")
		os.Exit(0)
	}

	view.PrintAddMemo(memo)
}
