package pgo

import "text/template"

var WhoAllianceTemplate = template.Must(template.New("whoAlliance").Parse(`{{ .Name }} ({{ .Acronym }})
https://politicsandwar.com/alliance/id={{ .ID }}
{{ .Score }}`))
