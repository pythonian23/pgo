package pgo

import "text/template"

var WhoAllianceTemplate = template.Must(template.New("whoAlliance").Parse(`# __{{.Name}}__ ({{.Acronym}})
*https://politicsandwar.com/alliance/id={{.ID}}*
**{{.Name}}** has a score of *{{.Score}}*, with *{{.Members}}* members and *{{.Applicants}}* applicants.
## Wars
- Involved in a total of *{{.TotalWars}}* wars.
- *{{.OffensiveWars}}* are offensive and *{{.DefensiveWars}}* are defensive.
- Raiding in *{{.OffensiveRaids}}* wars and are being raided in *{{.DefensiveRaids}}*.
- Winning roughly {{.WinningWarPercent}}% of their wars.`))
