package flags

import (
	flag "github.com/spf13/pflag"
)

func NewFlagSet(name string, flagSets ...*flag.FlagSet) *flag.FlagSet {
	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)
	flagSet.AddFlagSet(GeneralFlagSet)
	for _, fs := range flagSets {
		flagSet.AddFlagSet(fs)
	}
	return flagSet
}

func ReadArgs(fs *flag.FlagSet, args []string) (Args, error) {
	commonArgsLock.Lock()
	defer commonArgsLock.Unlock()
	err := fs.Parse(args)
	return commonArgs, err
}
