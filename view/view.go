package view

import (
	"fmt"

	"github.com/kcwebapply/imemo/data"
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
