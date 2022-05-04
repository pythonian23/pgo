package pgo

import "github.com/pythonian23/pgo/internal/flags"

// Command is a type for pgo's subcommands.
type Command struct {
	// Name is the name of the subcommand.
	Name        string
	// Action is the body of the subcommand.
	// It takes the provided arguments (excluding the subcommand name) and outputs the output string,
	// along with any error that may happen.
	Action      func([]string) (string, error)
	// The subcommand's description, for use in help messages.
	Description string
	// Aliases for the subcommand. These should also work when running the command.
	Aliases     []string
	flags.Flags
}

var commands = map[string]*Command{}
var aliases = map[string]string{}

// GetCommand gets a command by provided name.
// If the command does not exist, ok is false.
// Aliases are also searchable.
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

// AddCommand registers the given command, along with its aliases.
func AddCommand(command *Command) {
	commands[command.Name] = command
	for _, alias := range command.Aliases {
		aliases[alias] = command.Name
	}
}

// DeleteCommand deletes a command by name, and removes all aliases.
// Aliases do not work for the name.
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
