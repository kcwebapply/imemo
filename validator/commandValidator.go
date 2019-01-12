package validator

import (
	"errors"
	"fmt"

	"github.com/codegangsta/cli"
)

var argumentSizeMap = map[string]int{}

func init() {
	argumentSizeMap["a"] = 1
	argumentSizeMap["all"] = 1
	argumentSizeMap["s"] = 2
	argumentSizeMap["save"] = 2
	argumentSizeMap["d"] = 2
	argumentSizeMap["delete"] = 2
	argumentSizeMap["e"] = 2
	argumentSizeMap["edit"] = 2
	argumentSizeMap["c"] = 1
	argumentSizeMap["clear"] = 1
}

// ValidateArgument validate command args.
func ValidateArgument(c *cli.Context) error {
	var err error
	commandName := c.Args().Get(0)
	inputArgumentsSize := len(c.Args())
	correctSize, exists := argumentSizeMap[commandName]

	if !exists {
		message := fmt.Sprintf("command not found :%s", commandName)
		err = errors.New(message)
	}

	if inputArgumentsSize != correctSize {
		message := fmt.Sprintf("arguments size error.")
		err = errors.New(message)
	}
	return err
}
