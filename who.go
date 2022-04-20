package pgo

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var whoFlags *flag.FlagSet = flag.NewFlagSet("who", flag.ContinueOnError)
var help *bool = whoFlags.BoolP("help", "h", false, "Show help about this command.")
var alliance *bool = whoFlags.BoolP("alliance", "a", false, "Alliance mode switch. only links and IDs are allowed in alliance mode.")
var id *int = whoFlags.IntP("id", "i", 0, "The nation/alliance ID (just the number)")
var link *string = whoFlags.StringP("link", "l", "", "The nation/alliance link.")
var discordName *string = whoFlags.StringP("discord", "d", "", "The discord name")
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
	if *help {
		return HelpSubCommand(whoCmd), nil
	}
	if *id != 0 {
		request(whoQuery, whoData)
		return fmt.Sprintln(whoData), nil
	}
	return "not yet implemented", nil
}

func init() {
	whoCmd = &Command{"who", who, "get information about the nation/alliance", []string{}, whoFlags}
	whoFlags.SortFlags = false
	AddCommand(whoCmd)
}
