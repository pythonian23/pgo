package pgo

import "text/template"

var WhoAllianceTemplate = template.Must(template.New("whoAlliance").Parse(`# __{{.Data.Name}}__ ({{.Data.Acronym}})
*https://politicsandwar.com/alliance/id={{.Data.ID}}*
**{{.Data.Name}}** has a score of *{{.Data.Score}}*, with *{{.Members}}* members and *{{.Applicants}}* applicants.
## Wars
- Involved in a total of *{{.TotalWars}}* wars.
- *{{.OffensiveWars}}* are offensive and *{{.DefensiveWars}}* are defensive.
- Raiding in *{{.OffensiveRaids}}* wars and are being raided in *{{.DefensiveRaids}}*.
- Winning roughly {{.WinningWarPercent}}% of their wars.
## Treaties{{with .Data.Treaties}}{{range .}}
- **{{printf "%-12s" .TreatyType}}**:
_{{printf "%32s" .Alliance1.Name}}_ **--=#=--** _{{.Alliance2.Name}}_
{{end}}{{else}}
**NONE**{{end}}`))
