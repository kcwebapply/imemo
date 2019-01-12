package commands

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/codegangsta/cli"
	data "github.com/kcwebapply/imemo/data"
	view "github.com/kcwebapply/imemo/view"
)

var fileName = "data.txt"

//GetAllMemo shows all memodata
func GetAllMemo(c *cli.Context) {
	lines := readLines()
	view.PrintAllMemoMessage(lines)
}

// SaveMemo saves memodata
func SaveMemo(c *cli.Context) {
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	} else {
		log.Fatal("error")
	}
	newData := saveMemo(paramFirst)
	view.PrintSaveMessage(newData)
}

// DeleteMemo delete memodata
func DeleteMemo(c *cli.Context) {
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	} else {
		log.Fatal("error")
	}
	paramIntValue, err := strconv.Atoi(paramFirst)
	if err != nil {
		fmt.Println("argument error. first paramerter shouild be int-format")
	}
	deleteMemo(paramIntValue)
}

// EditMemo edit memodata
func EditMemo(c *cli.Context) {
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	} else {
		log.Fatal("argument error. please input memo id.")
	}
	paramIntValue, err := strconv.Atoi(paramFirst)
	if err != nil {
		fmt.Println("argument error. first paramerter shouild be int-format")
	}

	text, err := getTerminalInput()
	if err != nil {
		log.Fatal(err.Error())
	}
	editMemo(paramIntValue, text)
}

// ClearMemo delete all memo data
func ClearMemo(c *cli.Context) {
	clearMemo()
}

func readLines() []data.Data {
	var lines = []data.Data{}

	f, _ := os.Open(fileName)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data, err := data.ConvertToData(scanner.Text())
		if err != nil {
			continue
		}
		lines = append(lines, data)
	}
	lines = sortAndDeleteDuplication(lines)
	return lines
}

func saveMemo(text string) data.Data {
	lines := readLines()
	writer := getFileCleanWriter()
	defer writer.Flush()
	lineSize := len(lines)
	for _, data := range lines {
		writer.Write(([]byte)(data.String()))
	}

	var lastCOUNTER int
	if lineSize > 0 {
		lastCOUNTER = lines[len(lines)-1].Id + 1
	} else {
		lastCOUNTER = 1
	}
	newData := data.Data{Id: lastCOUNTER, Text: text}
	writer.Write(([]byte)(newData.String()))
	return newData
}

func deleteMemo(id int) {
	lines := readLines()
	writer := getFileCleanWriter()
	defer writer.Flush()
	counter := 1
	for _, data := range lines {
		if id != data.Id {
			data.Id = counter
			writer.Write(([]byte)(data.String()))
			writer.Flush()
			counter++
		} else {
			view.PrintDeleteMessage(data)
		}
	}
	writer.Write(([]byte)("\n"))
	writer.Flush()
}

func editMemo(id int, text string) {
	lines := readLines()
	writer := getFileCleanWriter()
	defer writer.Flush()
	counter := 1
	for _, data := range lines {
		if id != data.Id {
			data.Id = counter
			writer.Write(([]byte)(data.String()))
			writer.Flush()
			counter++
		} else {
			data.Id = counter
			data.Text = text
			writer.Write(([]byte)(data.String()))
			counter++
			view.PrintEditMessage(data)
		}
	}
	writer.Write(([]byte)("\n"))
	writer.Flush()
}

func clearMemo() {
	writer := getFileCleanWriter()
	defer writer.Flush()
	view.PrintClearMessage()
}

func getFileCleanWriter() *bufio.Writer {
	writeFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	writer := bufio.NewWriter(writeFile)
	return writer
}

func sortAndDeleteDuplication(datas []data.Data) []data.Data {
	dataSets := []data.Data{}
	idSets := make(map[int]struct{})

	for _, data := range datas {
		if _, ok := idSets[data.Id]; !ok {
			idSets[data.Id] = struct{}{}
			dataSets = append(dataSets, data)
		}
	}

	sort.Slice(dataSets, func(i, j int) bool {
		return dataSets[i].Id < dataSets[j].Id
	})
	return dataSets
}

func getTerminalInput() (string, error) {
	var err error
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		err = errors.New("get input error")
	}
	return scanner.Text(), err
}
