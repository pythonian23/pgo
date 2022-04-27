package flags

import (
	"sync"

	flag "github.com/spf13/pflag"
)

var generalFlags *flag.FlagSet = flag.NewFlagSet("general", flag.ContinueOnError)
var GeneralFlags = struct {
	*Flags
	Help *bool
	*sync.Mutex
}{
	Flags: SetToFlags(generalFlags),
	Help:  generalFlags.BoolP("help", "h", false, "Show help about this command."),
}
