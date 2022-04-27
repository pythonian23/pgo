package flags

import (
	"sync"

	flag "github.com/spf13/pflag"
)

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
