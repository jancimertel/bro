package _default

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"bitbucket.org/jmertel/bro/templates/types"
	"go/ast"
	"os"
	"sort"
	"text/template"
)

const tmpl = `
Packages:
	{{range $data := . }}
		-{{$data.Pkg}}
		{{range $func := .Funcs}}
			-{{$func}}
		{{end}}
	{{end}}
`

/*
	Default template can output provider data in simple form - formatted text file
*/
type defaultTemplate struct {
	provider analyserTypes.IProvider
}

func (d defaultTemplate) Serve(port string) error {
	panic("implement me")
}

func (d defaultTemplate) Build() error {
	var data []struct {
		Pkg   string
		Funcs []string
	}

	for _, pkg := range d.provider.GetPackages() {
		data = append(data, struct {
			Pkg   string
			Funcs []string
		}{
			pkg.Name,
			[]string{},
		})

		for _, file := range pkg.Files {
			for _, obj := range file.Scope.Objects {
				if obj.Kind == ast.Fun {
					data[len(data) - 1].Funcs = append(data[len(data) - 1].Funcs, obj.Name)
				}
			}
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Pkg < data[j].Pkg
	})

	t := template.Must(template.New("default").Parse(tmpl))
	outFile, err := os.Create("out")
	if err != nil {
		return err
	}
	return t.Execute(outFile, data)
}

func NewTemplate(provider analyserTypes.IProvider) types.ITemplater {
	return &defaultTemplate{
		provider,
	}
}
