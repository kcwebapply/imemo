package data

import (
	"errors"
	"strconv"
	"strings"
)

type Data struct {
	Id   int
	Text string
}

func (data *Data) String() string {
	return strconv.Itoa(data.Id) + " " + data.Text + "\n"
}

func ConvertToData(line string) (Data, error) {
	var data Data
	var err error
	datas := strings.Split(line, " ")
	text := generateText(datas)
	if len(datas) < 2 {
		err = errors.New("data parse error")
	} else {
		dataInt, _ := strconv.Atoi(datas[0])
		if dataInt == 0 {
			err = errors.New("malforlmed data error")
		}
		data = Data{Id: dataInt, Text: text}
	}
	return data, err
}

func generateText(tokens []string) string {
	var text = ""
	for i := 1; i < len(tokens); i++ {
		text += tokens[i] + " "
	}
	return text
}
