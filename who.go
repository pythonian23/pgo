package pgo

import (
	"fmt"
	"net/url"

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
		switch {
		case *arguments.Link != "":
			var link *url.URL
			link, err = url.Parse(*arguments.Link)
			if err != nil {
				break
			}
			_, err = fmt.Sscanf(link.EscapedPath(), "/alliance/id=%d", arguments.ID)
			if err != nil {
				break
			}
			fallthrough
		case *arguments.ID != 0:
			query = fmt.Sprintf("id:%d", *arguments.ID)
		case *arguments.Name != "":
			query = fmt.Sprintf("name:%q", *arguments.Name)
		default:
			err = fmt.Errorf("not enough information was provided")
			return
		}
		if err != nil {
			return
		}
		var alliance *api.Data
		alliance, err = request(fmt.Sprintf(whoAllianceQuery, query))
		if err == nil {
			out = fmt.Sprintln(alliance.Alliances.Data[0].Name)
		}
	} else {
		switch {
		case *arguments.Link != "":
			var link *url.URL
			link, err = url.Parse(*arguments.Link)
			if err != nil {
				break
			}
			_, err = fmt.Sscanf(link.EscapedPath(), "/nation/id=%d", arguments.ID)
			if err != nil {
				break
			}
			fallthrough
		case *arguments.ID != 0:
			query = fmt.Sprintf("id:%d", *arguments.ID)
		case *arguments.Name != "":
			query = fmt.Sprintf("nation_name:%q", *arguments.Name)
		case *arguments.Leader != "":
			query = fmt.Sprintf("leader_name:%q", *arguments.Leader)
		case *arguments.DiscordName != "":
			query = fmt.Sprintf("discord:%q", *arguments.DiscordName)
		default:
			err = fmt.Errorf("not enough information was provided")
		}
		if err != nil {
			return
		}
		var nation *api.Data
		nation, err = request(fmt.Sprintf(whoNationQuery, query))
		if err == nil {
			out = fmt.Sprintln(nation.Nations.Data[0].NationName)
		}
	}
	return
}

func init() {
	whoCmd = &Command{"who", who, "get information about the nation/alliance", []string{"whois"}, whoFlags}
	AddCommand(whoCmd)
}
