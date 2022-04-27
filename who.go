package pgo

import (
	"fmt"

	"github.com/pythonian23/pgo/flags"
	flag "github.com/spf13/pflag"
)

var whoFlags *flag.FlagSet = flags.NewFlagSet("who", flags.IdentityFlags)
var key *string = whoFlags.StringP("key", "k", "", "The API Key")
var whoCmd *Command

const whoQuery = "{nations(id:6) {data{nation_name}}}"

var whoData = struct {
	Errors []struct{ Message string }
	Data   struct {
		Nations []struct {
			NationName string `json:"nation_name"`
		}
	}
}{}

func who(args []string) (string, error) {
	err := whoFlags.Parse(args)
	if err != nil {
		return "", err
	}
	Key = *key
	if *flags.GeneralFlags.Help {
		return HelpSubCommand(whoCmd), nil
	}
	if *flags.IdentityFlags.ID != 0 {
		request(whoQuery, whoData)
		return fmt.Sprintln(whoData), nil
	}
	return "not yet implemented", nil
}

func init() {
	whoCmd = &Command{"who", who, "get information about the nation/alliance", []string{"whois"}, whoFlags}
	whoFlags.SortFlags = false
	AddCommand(whoCmd)
}
