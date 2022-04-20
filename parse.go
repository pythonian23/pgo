package pgo

func Parse(args []string) string {
	if len(args) < 2 {
		return "usage: pgo <SUBCOMMAND> [OPTIONS]"
	}
	command, ok := GetCommand(args[1])
	if !ok {
		return NoCommandHandler(args[1])
	}
	out, err := command.Action(args[2:])
	if err != nil {
		return ErrorHandler(err)
	}
	return out
}
