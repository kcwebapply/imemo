package view

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kcwebapply/imemo/data"
	"github.com/kcwebapply/imemo/domain/model"
	util "github.com/kcwebapply/imemo/util"
)

var dotLine = "------------------------------------------------------------------------------------------"

// PrintAllMemoMessage is function of printing message when showing all memo.
/*func PrintAllMemoMessage(datas []data.Data) {
	fmt.Println(dotLine)

	for _, data := range datas {
		printBody(data)
	}
	fmt.Println(dotLine + "\x1b[0m")
}*/

func PrintAllMemo(memos []model.Memo) {
	fmt.Println(dotLine)

	for _, memo := range memos {
		printBody(memo)
	}
	fmt.Println(dotLine + "\x1b[0m")
}

// PrintSaveMessage is function of printing message when saving memo.
func PrintSaveMessage(data data.Data) {
	fmt.Println(dotLine)
	//printBody(data)
	fmt.Println(dotLine)
	//fmt.Println("\x1b[1;94mmemo saved!\x1b[0m")
	fmt.Println("\x1b[1m\x1b[38;5;39mmemo saved!\x1b[0m")
}

// PrintDeleteMessage is function of printing message when deleting memo.
func PrintDeleteMessage(data data.Data) {
	fmt.Println(dotLine)
	//printBody(data)
	fmt.Println(dotLine)
	//fmt.Println("\x1b[1;36mmemo deleted!\x1b[0m")
	fmt.Println("\x1b[1m\x1b[38;5;211mmemo deleted!\x1b[0m")
}

// PrintEditMessage is function of printing message when editing memo.
func PrintEditMessage(data data.Data) {
	fmt.Println(dotLine)
	//printBody(data)
	fmt.Println(dotLine)
	//fmt.Println("\x1b[1;33;21mmemo edited!\x1b[0m")
	fmt.Println("\x1b[1m\x1b[38;5;214mmemo edited!\x1b[0m")
}

// PrintClearMessage is function of printing message when clearing memo.
func PrintClearMessage(size int) {
	message := fmt.Sprintf("\x1b[1;36mclear %d memo!\x1b[0m", size)
	fmt.Println(message)
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
