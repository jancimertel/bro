package text

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"bitbucket.org/jmertel/bro/templates/types"
	"encoding/json"
	"io/ioutil"
	"os"
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
	Default template can output provider data in json, then show it rendered in web template
*/
type defaultTemplate struct {
	provider analyserTypes.IProvider
}

func (d defaultTemplate) Serve(port string) error {
	panic("implement me")
}

func (d *defaultTemplate) Build() error {
	if err := d.buildJsonData(); err != nil {
		return err
	}

	return nil
}

func (d *defaultTemplate) buildJsonData() error {
	bytes, err := json.Marshal(d.provider.Dump())
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("out.json", bytes, os.ModePerm); err != nil {
		return err
	}

	return nil
}
func NewTemplate(provider analyserTypes.IProvider) types.ITemplater {
	return &defaultTemplate{
		provider,
	}
}
