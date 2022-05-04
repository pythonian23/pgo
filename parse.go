package pgo

import flag "github.com/spf13/pflag"

// Parse is the entrypoint for running pgo.
// It requires arguments in an os.Args-like style and outputs a string.
func Parse(args []string) string {
	if len(args) < 2 {
		return "usage: pgo <SUBCOMMAND> [OPTIONS]"
	}
	command, ok := GetCommand(args[1])
	if !ok {
		return NoCommandHandler(args[1])
	}
	out, err := command.Action(args[2:])
	if err != nil && err != flag.ErrHelp {
		return ErrorHandler(err)
	}
	return out
}
