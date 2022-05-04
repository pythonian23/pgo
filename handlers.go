package pgo

import (
	"fmt"
)

// ErrorHandler is the function for handling errors.
// It should take an error and output a string in the desired format.
var ErrorHandler func(error) string = func(err error) string {
	return err.Error()
}

// NoCommandHandler is the function for handling commands that were requested but do not exist.
// It should take the requested command name as a string and output a string.
var NoCommandHandler func(string) string = func(command string) string {
	return fmt.Sprintf("Command %s does not exist.", command)
}
