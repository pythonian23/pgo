package pgo

import flag "github.com/spf13/pflag"

type Command struct {
	Name string
	Action func([]string) (string, error)
	Description string
	*flag.FlagSet
}

var Commands map[string]Command = map[string]Command{}

func AddCommand(command Command) {
	Commands[command.Name] = command
}
