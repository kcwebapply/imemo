package view

import (
	"fmt"
	"strconv"
	"strings"

	data "github.com/kcwebapply/imemo/data"
	util "github.com/kcwebapply/imemo/util"
)

var dotLine = "------------------------------------------------------------------------------------------"

// PrintAllMemoMessage is function of printing message when showing all memo.
func PrintAllMemoMessage(datas []data.Data) {
	fmt.Println(dotLine)

	for _, data := range datas {
		printBody(data)
	}
	fmt.Println(dotLine + "\x1b[0m")
}

// PrintSaveMessage is function of printing message when saving memo.
func PrintSaveMessage(data data.Data) {
	fmt.Println(dotLine)
	printBody(data)
	fmt.Println(dotLine)
	fmt.Println("\x1b[1;32mmemo saved!\x1b[0m")
}

// PrintDeleteMessage is function of printing message when deleting memo.
func PrintDeleteMessage(data data.Data) {
	fmt.Println(dotLine)
	printBody(data)
	fmt.Println(dotLine)
	fmt.Println("\x1b[1;36mmemo deleted!\x1b[0m")
}

// PrintEditMessage is function of printing message when editing memo.
func PrintEditMessage(data data.Data) {
	fmt.Println(dotLine)
	printBody(data)
	fmt.Println(dotLine)
	fmt.Println("\x1b[1;33mmemo edited!\x1b[0m")
}

// PrintClearMessage is function of printing message when clearing memo.
func PrintClearMessage(size int) {
	message := fmt.Sprintf("\x1b[1;36mclear %d memo!\x1b[0m", size)
	fmt.Println(message)
}

func printBody(data data.Data) {
	echo := ""
	echo += "|"
	if data.Id < 10 {
		echo += " "
	}

	echo += strconv.Itoa(data.Id)
	echo += "| "
	echo += data.Text
	textLength := util.TextCounter(echo)
	space := len(dotLine) - textLength - 1
	echo += strings.Repeat(" ", space)
	echo += "|"
	fmt.Println(echo)
}
