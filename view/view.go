package view

import (
	"fmt"
	"strconv"
	"strings"

	data "github.com/kcwebapply/imemo/data"
)

var dotLine = "----------------------------------------------------------------------------------"

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

func printBody(data data.Data) {
	echo := ""
	echo += "|"
	if data.Id < 10 {
		echo += " "
	}

	echo += strconv.Itoa(data.Id)
	echo += "| "
	echo += data.Text
	textLength := textCounter(echo)
	space := len(dotLine) - textLength - 1
	echo += strings.Repeat(" ", space)
	echo += "|"
	fmt.Println(echo)
}

func textCounter(text string) int {
	textCounter := 0
	befPos := 0
	for pos, _ := range text {
		// this check is where character is Japanase or not.
		if pos-befPos == 3 {
			textCounter += 2 // to treat japanese character as 2byte.
			befPos = pos
		} else {
			textCounter += 1
			befPos = pos
		}
	}
	return textCounter
}
