package builder

import (
	"html/template"
	"io"
)

const templateContent = `
public static class Builder {
{{range .Variables}}
    private {{.Type}} {{.Name}};{{end}}
{{range .Variables}}
    public Builder {{.Name}}(final {{.Type}} {{.Name}}) {
        this.{{.Name}} = {{.Name}};
        return this;
    }
{{end}}
    public {{.ClassName}} build() {
        return new {{.ClassName}}(this);
    }
}
`

type class struct {
	Variables []variable
	ClassName string
}

type variable struct {
	Type template.HTML
	Name string
}

// Generate generates the builder class code from class name
// and variables. It writes the output to w.
func Generate(w io.Writer, className string, vars []Variable) error {
	tmpl := template.Must(template.New("builder").Parse(templateContent))

	var vv []variable
	for _, v := range vars {
		vv = append(vv, variable{
			Type: template.HTML(v.Type()),
			Name: v.Name(),
		})
	}

	return tmpl.Execute(w, class{
		ClassName: className,
		Variables: vv,
	})
}
