package commands

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/codegangsta/cli"
	data "github.com/kcwebapply/imemo/data"
	util "github.com/kcwebapply/imemo/util"
	view "github.com/kcwebapply/imemo/view"
)

var fileName = ""

var maxTextSize = 60

func GetCommandMap() []cli.Command {
	var commands = []cli.Command{
		{
			Name:    "ls",
			Aliases: []string{},
			Usage:   "list memo. ",
			Action:  Ls,
		},
		{
			Name:    "add",
			Aliases: []string{"add"},
			Usage:   "Save memo.",
			Action:  Add,
		},
		{
			Name:    "rm",
			Aliases: []string{"r"},
			Usage:   "Delete memo that specified by id.",
			Action:  Rm,
		},
		{
			Name:    "edit",
			Aliases: []string{"e"},
			Usage:   "Edit memo that specified by id.",
			Action:  EditMemo,
		},
		{
			Name:    "clear",
			Aliases: []string{"c"},
			Usage:   "clear (delete) all memo data.",
			Action:  ClearMemo,
		},
	}
	return commands
}

// EditMemo edit memodata
func EditMemo(c *cli.Context) {
	paramIntValue, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fmt.Println("argument error. first paramerter should be int-format")
	}

	text, err := util.GetTerminalInput()
	if err != nil {
		fmt.Println("input error happened.")
		os.Exit(0)
	}

	if util.TextCounter(text) > maxTextSize {
		fmt.Printf("argument error. text size should be under %d.\n", maxTextSize)
		os.Exit(0)
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
	lines := readLines()
	writer := getFileCleanWriter()
	defer writer.Flush()
	view.PrintClearMessage(len(lines))
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
