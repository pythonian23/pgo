package pgo

import (
	"fmt"

	"github.com/pythonian23/pgo/internal/flags"
)

var whoFlags = flags.NewFlags("who", flags.IdentityFlagSet, flags.APIFlagSet)
var whoCmd *Command

const whoQuery = "{nations(id:6) {data{nation_name}}}"

func who(args []string) (string, error) {
	arguments, err := flags.ReadArgs(whoFlags, args)
	if err != nil {
		return "", err
	}
	Key = *arguments.APIKey
	if *arguments.Help {
		return HelpSubCommand(whoCmd), nil
	}
	if *arguments.ID != 0 {
		data, err := request(whoQuery)
		return fmt.Sprintln(data.Nations), err
	}
	return "not yet implemented", nil
}

func init() {
	whoCmd = &Command{"who", who, "get information about the nation/alliance", []string{"whois"}, whoFlags}
	AddCommand(whoCmd)
}
