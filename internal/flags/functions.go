package flags

import (
	flag "github.com/spf13/pflag"
)

// NewFlags composes a new Flags object from the provided FlagSets.
// GeneralFlagSet is added by default.
func NewFlags(name string, flagSets ...*flag.FlagSet) Flags {
	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)
	flagSet.SortFlags = false
	flagSet.AddFlagSet(GeneralFlagSet)
	for _, fs := range flagSets {
		flagSet.AddFlagSet(fs)
	}
	return Flags{flagSet}
}

func ReadArgs(flags Flags, args []string) (Args, error) {
	commonArgsLock.Lock()
	defer commonArgsLock.Unlock()
	err := flags.Parse(args)
	return commonArgs, err
}
