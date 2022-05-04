package flags

import (
	"sync"

	flag "github.com/spf13/pflag"
)

var GeneralFlagSet = flag.NewFlagSet("general", flag.ContinueOnError)
var IdentityFlagSet = flag.NewFlagSet("identity", flag.ContinueOnError)
var APIFlagSet = flag.NewFlagSet("api", flag.ContinueOnError)

var commonArgsLock = sync.Mutex{}
var commonArgs = Args{
	generalArgs: &generalArgs{
		Help: GeneralFlagSet.BoolP("help", "h", false, "Show help about this command."),
	},

	identityArgs: &identityArgs{
		Alliance:    IdentityFlagSet.BoolP("alliance", "a", false, "Alliance mode switch. only links and IDs are allowed in alliance mode."),
		ID:          IdentityFlagSet.IntP("id", "i", 0, "The nation/alliance ID (just the number)"),
		Link:        IdentityFlagSet.StringP("link", "l", "", "The nation/alliance link."),
		DiscordName: IdentityFlagSet.StringP("discord", "d", "", "The discord name"),
	},
	apiArgs: &apiArgs{
		APIKey: APIFlagSet.StringP("key", "K", "", "The API key to be used for accessing the API"),
	},
}
