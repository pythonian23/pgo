package pgo

import "github.com/jessevdk/go-flags"

type whoArgs struct {
	Discord string `short:"d" long:"discord"`
	GlobalArgs
}

// Who returns information about the specified nation.
func Who(args ...string) (msg string, ok bool) {
	opt := flags.
	parsed, err := flags.ParseArgs(&opt, args)
}
