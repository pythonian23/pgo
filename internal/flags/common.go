package flags

import (
	"sync"

	flag "github.com/spf13/pflag"
)

// GeneralFlagSet contains flags that should be in every command.
var GeneralFlagSet = flag.NewFlagSet("general", flag.ContinueOnError)

// IdentityFlagSet contains flags for identifying a desired nation or alliance.
var IdentityFlagSet = flag.NewFlagSet("identity", flag.ContinueOnError)

// APIFlagSet contains flags for accessing the API.
var APIFlagSet = flag.NewFlagSet("api", flag.ContinueOnError)

var commonArgsLock = sync.Mutex{}
var commonArgs = Args{
	generalArgs: generalArgs{
		Help: GeneralFlagSet.BoolP("help", "h", false, "Show help about this command."),
	},

	identityArgs: identityArgs{
		Alliance:    IdentityFlagSet.BoolP("alliance", "a", false, "Alliance mode switch. only links and IDs are allowed in alliance mode."),
		ID:          IdentityFlagSet.IntP("id", "i", 0, "The nation/alliance ID (just the number)."),
		Link:        IdentityFlagSet.StringP("link", "l", "", "The nation/alliance link."),
		Name:        IdentityFlagSet.StringP("name", "n", "", "The nation/alliance name."),
		Leader:      IdentityFlagSet.StringP("leader", "N", "", "The nation's leader name."),
		DiscordName: IdentityFlagSet.StringP("discord", "d", "", "The discord name."),
	},
	apiArgs: apiArgs{
		APIKey: APIFlagSet.StringP("key", "K", "", "The API key to be used for accessing the API"),
	},
}
