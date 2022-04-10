package pgo

import (
	"fmt"
)

var ErrorHandler func(error) string = func(err error) string {
	return err.Error()
}

var NoCommandHandler func(string) string = func(command string) string {
	return fmt.Sprintf("Command %s does not exist.", command)
}
