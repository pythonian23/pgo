package flags

import (
	flag "github.com/spf13/pflag"
)

// Flags is a wrapping type for pflag.FlagSet to allow repeated use of the same flagset.
type Flags struct {
	*flag.FlagSet
}

// Parse resets the variables before parsing as usual.
func (f Flags) Parse(args []string) error {
	f.Reset()
	return f.FlagSet.Parse(args)
}

// Reset resets the variables.
func (f Flags) Reset() {
	f.Visit(func(flag *flag.Flag) {
		flag.Changed = false
		if err := flag.Value.Set(flag.DefValue); err != nil {
			panic(err)
		}
	})
}
