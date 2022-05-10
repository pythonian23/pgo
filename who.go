package pgo

import (
	"fmt"

	"github.com/pythonian23/pgo/internal/api"
	"github.com/pythonian23/pgo/internal/flags"
)

var whoFlags = flags.NewFlags("who", flags.IdentityFlagSet, flags.APIFlagSet)
var whoCmd *Command

const whoNationQuery = "{nations(%s){data{nation_name}}}"
const whoAllianceQuery = "{alliances(%s){data{name}}}"

func who(args []string) (out string, err error) {
	arguments, err := flags.ReadArgs(whoFlags, args)
	if err != nil {
		return
	}
	if *arguments.Help {
		return HelpSubCommand(whoCmd), nil
	}
	Key = *arguments.APIKey
	var query string
	if *arguments.Alliance {
		var alliance *api.Data
		switch {
		case *arguments.ID != 0:
			query = fmt.Sprintf("id:%d", *arguments.ID)
		default:
			err = fmt.Errorf("not enough information was provided")
			return
		}
		alliance, err = request(fmt.Sprintf(whoAllianceQuery, query))
		if err == nil {
			out = fmt.Sprintln(alliance.Alliances.Data[0].Name)
		}
	} else {
		switch {
		default:
			err = fmt.Errorf("not enough information was provided")
		}
	}

	return
}

func init() {
	whoCmd = &Command{"who", who, "get information about the nation/alliance", []string{"whois"}, whoFlags}
	AddCommand(whoCmd)
}
