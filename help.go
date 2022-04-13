package pgo

import (
	flag "github.com/spf13/pflag"
	"text/template"
	"bytes"
)

var HelpTemplate *template.Template = template.Must(template.New("help").Funcs(template.FuncMap{"usages": func(fs *flag.FlagSet)string{return fs.FlagUsages()}}).Parse(`{{ .Name }} - {{ .Description }}
ALIASES
ARGUMENTS
{{ .FlagSet.FlagUsages }}`))

func HelpSubCommand(command Command) string {
	buf := &bytes.Buffer{}
	err := HelpTemplate.Execute(buf, command)
	if err != nil {
		return SeriousErrorHandler(err)
	}
	return buf.String()
}
