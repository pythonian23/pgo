package flags

import (
	"sync"

	flag "github.com/spf13/pflag"
)

type Flagger interface {
	FlagSet() *flag.FlagSet
	Lock()
	Unlock()
}

// NewFlagSet returns a new FlagSet with the given name.
// The FlagSetters in flags are added,
// and GeneralFlags is added by default.
func NewFlagSet(name string, flags ...Flagger) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	for _, flags := range flags {
		fs.AddFlagSet(flags.FlagSet())
	}
	return fs
}

type Flags flag.FlagSet

func SetToFlags(flagSet *flag.FlagSet) *Flags {
	return (*Flags)(flagSet)
}

func (f *Flags) FlagSet() *flag.FlagSet {
	return (*flag.FlagSet)(f)
}

// Parse runs Reset, then parses the given arguments.
func (f *Flags) Parse(args []string) error {
	return f.FlagSet().Parse(args)
}

// Reset resets all set values in the Flag.
// Panics if any flag is unable to be set to its default value.
func (f *Flags) Reset() {
	f.FlagSet().Visit(func(f *flag.Flag) {
		f.Changed = false
		if err := f.Value.Set(f.DefValue); err != nil {
			panic(err)
		}
	})
}

var generalFlags *flag.FlagSet = flag.NewFlagSet("general", flag.ContinueOnError)
var GeneralFlags = struct {
	*Flags
	Help *bool
	*sync.Mutex
}{
	Flags: SetToFlags(generalFlags),
	Help:  generalFlags.BoolP("help", "h", false, "Show help about this command."),
}

var identityFlags *flag.FlagSet = flag.NewFlagSet("identity", flag.ContinueOnError)
var IdentityFlags = struct {
	*Flags
	Alliance    *bool
	ID          *int
	Link        *string
	DiscordName *string
	*sync.Mutex
}{
	Flags:       SetToFlags(identityFlags),
	Alliance:    identityFlags.BoolP("alliance", "a", false, "Alliance mode switch. only links and IDs are allowed in alliance mode."),
	ID:          identityFlags.IntP("id", "i", 0, "The nation/alliance ID (just the number)"),
	Link:        identityFlags.StringP("link", "l", "", "The nation/alliance link."),
	DiscordName: identityFlags.StringP("discord", "d", "", "The discord name"),
}
