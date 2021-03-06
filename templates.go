package pgo

import "text/template"

var WhoAllianceTemplate = template.Must(template.New("whoAlliance").Parse(`# __{{.Data.Name}}__ ({{.Data.Acronym}})
*https://politicsandwar.com/alliance/id={{.Data.ID}}*
## General Information
- Total score: *{{.Data.Score}}*
- *{{.Cities}}* cities in total.
- *{{.Members}}* members and *{{.Applicants}}* applicants.
- Color: **{{.Data.Color}}**
## Wars
- Involved in a total of *{{.TotalWars}}* wars.
- *{{.OffensiveWars}}* are offensive and *{{.DefensiveWars}}* are defensive.
- Raiding in *{{.OffensiveRaids}}* wars and are being raided in *{{.DefensiveRaids}}*.
- Winning roughly {{.WinningWarPercent}}% of their wars.
## Treaties{{with .Data.Treaties}}{{range .}}
- **{{printf "%-12s" .TreatyType}}**: {{$oneIsAlliance := eq .Alliance1.Name $.Data.Name}}{{if $oneIsAlliance}}*{{end}}*{{printf "%24s" .Alliance1.Name}}*{{if $oneIsAlliance}}*{{end}} --=#=-- {{if not $oneIsAlliance}}*{{end}}*{{.Alliance2.Name}}*{{if not $oneIsAlliance}}*{{end}}
{{end}}{{else}}
**NONE**{{end}}`))
