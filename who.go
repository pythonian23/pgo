package pgo

import flag "github.com/spf13/pflag"

var whoFlags *flag.FlagSet = flag.NewFlagSet("who", flag.ContinueOnError)
var help *bool = whoFlags.BoolP("help", "h", false, "Show help about this command.")
var alliance *bool = whoFlags.BoolP("alliance", "a", false, "Alliance mode switch. only links and IDs are allowed in alliance mode.")
var id  *int = whoFlags.IntP("id", "i", -1,  "The nation/alliance ID (just the number)")
var link *string = whoFlags.StringP("link", "l", "", "The nation/alliance link.")
var discordName *string = whoFlags.StringP("discord", "d", "", "The discord name")

func who(args []string) (string, error) {
	err := whoFlags.Parse(args)
	if err != nil {
		return "", err
	}
	if *help {
		return HelpSubCommand(Commands["who"]), nil
	}
	return "not yet implemented", nil
}

func init() {
	whoFlags.SortFlags =  false
	AddCommand(Command{"who", who, "get information about the nation/alliance", whoFlags})
}
