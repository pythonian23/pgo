package pgo

import (
	"bytes"
	"text/template"

	flag "github.com/spf13/pflag"
)

var HelpTemplate *template.Template = template.Must(template.New("help").Funcs(template.FuncMap{"usages": func(fs *flag.FlagSet) string { return fs.FlagUsages() }}).Parse(`{{ .Name }} - {{ .Description }}
ALIASES{{ range .Aliases }}
{{ . }}{{end}}
ARGUMENTS
{{ .FlagSet.FlagUsages }}`))

func HelpSubCommand(command *Command) string {
	buf := &bytes.Buffer{}
	err := HelpTemplate.Execute(buf, command)
	if err != nil {
		return SeriousErrorHandler(err)
	}
	return buf.String()
}
