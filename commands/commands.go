package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/codegangsta/cli"
	config "github.com/kcwebapply/imemo/config"
	data "github.com/kcwebapply/imemo/data"
)

var fileName = config.GetConfig().Sys.FileName

//GetAllMemo shows all memodata
func GetAllMemo(c *cli.Context) {
	lines := readLines()
	for _, data := range lines {
		fmt.Print(data.String())
	}
}

// SaveMemo saves memodata
func SaveMemo(c *cli.Context) {
	var paramFirst = ""
	if len(c.Args()) > 0 {
		paramFirst = c.Args().First()
	} else {
		log.Fatal("error")
	}
	saveMemo(paramFirst)
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

func saveMemo(text string) {
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
	fmt.Println("memo saved!")
	writer.Write(([]byte)(newData.String()))
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
		}
	}

	writer.Write(([]byte)("\n"))
	writer.Flush()
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
		//idSets[data.Id] = struct{}{}
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
