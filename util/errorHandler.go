package util

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Printf(err.Error())
	os.Exit(0)
}
