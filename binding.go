package pgo

type Command struct {
	Action      func(args []string) (string, error)
	Description string
}

var Commands map[string]Command
