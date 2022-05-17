package pgo

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/pythonian23/pgo/internal/api"
	"github.com/pythonian23/pgo/internal/flags"
)

var whoFlags = flags.NewFlags("who", flags.IdentityFlagSet, flags.APIFlagSet)
var whoCmd *Command

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
				return
			}
			_, err = fmt.Sscanf(link.EscapedPath(), "/alliance/id=%d", arguments.ID)
			if err != nil {
				return
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
		var data *api.Data
		data, err = request(fmt.Sprintf(whoAllianceQuery, query))
		if len(data.Alliances.Data) == 0 {
			err = fmt.Errorf("the requested alliance was not found")
		}
		if err != nil {
			return
		}
		alliance := struct {
			Data                                    *api.Alliance
			Members, Applicants                     int
			TotalWars, OffensiveWars, DefensiveWars int
			OffensiveRaids, DefensiveRaids          int
			winningWars                             int
			WinningWarPercent                       int
		}{Data: &data.Alliances.Data[0]}
		for _, nation := range alliance.Data.Nations {
			if nation.AlliancePosition == api.PositionApplicant {
				alliance.Applicants++
				continue
			}
			alliance.Members++
			for _, war := range nation.Wars {
				offensive := war.AttAllianceID == alliance.Data.ID
				raid := war.WarType == api.WarRaid
				if offensive {
					alliance.OffensiveWars++
					if war.AttResistance > war.DefResistance {
						alliance.winningWars++
					}
					if raid {
						alliance.OffensiveRaids++
					}
				} else {
					alliance.DefensiveWars++
					if war.AttResistance < war.DefResistance {
						alliance.winningWars++
					}
					if raid {
						alliance.DefensiveRaids++
					}
				}
			}
		}
		alliance.TotalWars = alliance.OffensiveWars + alliance.DefensiveWars
		alliance.WinningWarPercent = 100 * alliance.winningWars / alliance.TotalWars
		buf := &bytes.Buffer{}
		err = WhoAllianceTemplate.Execute(buf, alliance)
		if err != nil {
			return
		}
		out = buf.String()
	} else {
		switch {
		case *arguments.Link != "":
			var link *url.URL
			link, err = url.Parse(*arguments.Link)
			if err != nil {
				return
			}
			_, err = fmt.Sscanf(link.EscapedPath(), "/nation/id=%d", arguments.ID)
			if err != nil {
				return
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
		var data *api.Data
		data, err = request(fmt.Sprintf(whoNationQuery, query))
		if len(data.Nations.Data) == 0 {
			err = fmt.Errorf("the requested nation was not found")
		}
		if err == nil {
			out = fmt.Sprintln(data.Nations.Data[0].NationName)
		}
	}
	return
}

func init() {
	whoCmd = &Command{"who", who, "get information about the nation/alliance", []string{"whois"}, whoFlags}
	AddCommand(whoCmd)
}
