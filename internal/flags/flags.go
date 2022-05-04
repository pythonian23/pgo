package flags

import (
	flag "github.com/spf13/pflag"
)

type Flags struct {
	*flag.FlagSet
}

func (f Flags) Parse(args []string) error {
	f.Reset()
	return f.FlagSet.Parse(args)
}

func (f Flags) Reset() {
	f.Visit(func(flag *flag.Flag) {
		flag.Changed = false
		if err := flag.Value.Set(flag.DefValue); err != nil {
			panic(err)
		}
	})
}
