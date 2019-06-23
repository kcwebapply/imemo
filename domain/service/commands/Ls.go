package commands

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/imemo/domain/model"
	"github.com/kcwebapply/imemo/domain/repository"
	"github.com/kcwebapply/imemo/util"
	"github.com/kcwebapply/imemo/view"
)

func Ls(c *cli.Context) {
	var memolist []model.Memo

	var err error

	memolist, err = repository.GetMemos()

	if err != nil {
		util.HandleError(err)
	}

	categoryList, err := repository.GetCategories()

	if err != nil {
		util.HandleError(err)
	}

	memoCategoryMap := clustringByCategory(memolist)
	categoryIDMap := mapCategoryWithId(categoryList)
	view.PrintAllMemos(memoCategoryMap, categoryIDMap)
}

func clustringByCategory(memolist []model.Memo) map[int][]model.Memo {
	var memoCategoryMap = map[int][]model.Memo{}

	for _, memo := range memolist {
		// get categoryId
		memosCategoryID := memo.Categoryid
		// get all memolist thah have same category id
		var memolistByCategory, containCheck = memoCategoryMap[memosCategoryID]
		// if no memolist found yet, then mapping new hashmap
		if !containCheck {
			memolistByCategory = []model.Memo{}
		}

		memoCategoryMap[memosCategoryID] = append(memolistByCategory, memo)
	}

	return memoCategoryMap
}

func mapCategoryWithId(categoryList []model.Category) map[int]model.Category {
	var categoryIDMap = map[int]model.Category{}

	for _, category := range categoryList {
		categoryIDMap[category.ID] = category
	}

	return categoryIDMap
}
