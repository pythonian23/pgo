package pgo

import (
	"bytes"
	"text/template"

	flag "github.com/spf13/pflag"
)

// HelpTemplate is the template for help messages.
var HelpTemplate = template.Must(template.New("help").Funcs(template.FuncMap{"usages": func(fs *flag.FlagSet) string { return fs.FlagUsages() }}).Parse(`{{ .Name }} - {{ .Description }}
ALIASES{{ range .Aliases }}
{{ . }}{{end}}
ARGUMENTS
{{ .FlagSet.FlagUsages }}`))

// HelpSubCommand provides help for the provided subcommand using HelpTemplate.
func HelpSubCommand(command *Command) string {
	buf := &bytes.Buffer{}
	err := HelpTemplate.Execute(buf, command)
	if err != nil {
		return ErrorHandler(err)
	}
	return buf.String()
}
