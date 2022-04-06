package pgo

// CommandFunc is the standard function for all commands.
type CommandFunc func(...string) (string, bool)

type Command struct {
	CommandFunc
	Description string
}

// Commands links each command to a string that it can be called with in implementations.
// User-defined commands can be added and preexisting commands can be removed.
// Any implementation of pgo should use Commands to determine the commands used.
var Commands map[string]Command = map[string]Command{
	"who": Command{Who, "Returns information about the specified nation"},
}
