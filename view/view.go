package view

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kcwebapply/bm/util"
	"github.com/kcwebapply/imemo/domain/model"
)

var dotLine = "------------------------------------------------------------------------------------------"

func PrintAllMemo(memos []model.Memo) {
	fmt.Println(dotLine)

	for _, memo := range memos {
		printBody(memo)
	}

	fmt.Println(dotLine + "\x1b[0m")
}

func PrintAllMemos(memoMapByCategory map[int][]model.Memo, categoryIDMap map[int]model.Category) {

	// loop for each category
	for categoryID, memolist := range memoMapByCategory {
		category := categoryIDMap[categoryID]
		fmt.Printf("\x1b[38;5;204m[ %s ]\n\x1b[0m", category.Name)
		fmt.Println(dotLine)
		for _, memo := range memolist {
			printBody(memo)
		}
		fmt.Println(dotLine + "\x1b[0m")

	}

}

// PrintSaveMessage is function of printing message when saving memo.
func PrintAddMemo(memo model.Memo) {
	//fmt.Println("\x1b[1;94mmemo saved!\x1b[0m")
	fmt.Printf("\x1b[38;5;081m>> %s\n", memo.Memo)
	fmt.Println("\x1b[0mmemo saved!")
}

func PrintDeleteMemo(memo model.Memo) {
	//fmt.Println("\x1b[1;94mmemo saved!\x1b[0m")
	fmt.Printf("\x1b[38;5;173m>> %s\n", memo.Memo)
	fmt.Println("\x1b[0mmemo deleted!")
}

func PrintClearMemo(memo model.Memo) {
	fmt.Println("\x1b[0m all memo deleted!")
}

func printBody(memo model.Memo) {
	echo := ""
	echo += "|"
	if memo.ID < 10 {
		echo += " "
	}

	echo += strconv.Itoa(memo.ID)
	echo += "| "
	echo += memo.Memo
	textLength := util.TextCounter(echo)
	space := len(dotLine) - textLength - 1
	echo += strings.Repeat(" ", space)
	echo += "|"
	fmt.Println(echo)
}
