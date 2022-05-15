package pgo

import "text/template"

var WhoAllianceTemplate = template.Must(template.New("whoAlliance").Parse(`__**{{.Name}}**__ ({{.Acronym}})
*https://politicsandwar.com/alliance/id={{.ID}}*
**{{.Name}}** has a score of *{{.Score}}*, with *{{.Members}}* members and *{{.Applicants}}* applicants. They're involved in a total of *{{.TotalWars}}* wars. *{{.OffensiveWars}}* of those are offensive and *{{.DefensiveWars}}* are defensive. Of those, they are raiding in *{{.OffensiveRaids}}* wars and are being raided in *{{.DefensiveRaids}}*. They have a total of *{{.Resistance}}* resistance and their opponents have a total of *{{.OpponentResistance}}*.`))
