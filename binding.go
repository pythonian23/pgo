package pgo

import "github.com/pythonian23/pgo/internal/flags"

type Command struct {
	Name        string
	Action      func([]string) (string, error)
	Description string
	Aliases     []string
	flags.Flags
}

var commands map[string]*Command = map[string]*Command{}
var aliases map[string]string = map[string]string{}

func GetCommand(name string) (cmd *Command, ok bool) {
	for alias, command := range aliases {
		if alias == name {
			cmd, ok = commands[command]
			return
		}
	}
	cmd, ok = commands[name]
	return
}

func AddCommand(command *Command) {
	commands[command.Name] = command
	for _, alias := range command.Aliases {
		aliases[alias] = command.Name
	}
}

func DeleteCommand(name string) {
	command, ok := commands[name]
	if !ok {
		return
	}
	for _, alias := range command.Aliases {
		delete(aliases, alias)
	}
	delete(commands, name)
}
