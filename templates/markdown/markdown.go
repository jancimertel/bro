package markdown

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"bitbucket.org/jmertel/bro/templates"
	"bitbucket.org/jmertel/bro/templates/types"
	"go/ast"
	"os"
	"path"
	"text/template"
)

const tmpl = `
# Package {{.Pkg}}({{.PkgPath}})
{{.PkgComment}}
## Functions
{{range $func := .Funcs}}` +
	`-{{$func.Name}}({{$func.Signature.Text}})
{{$func.Comment}}
{{end}}
`

var mdTemplate *template.Template

/*
	Markdown template will generate folder structure with packages info in markdown files.
	These could be server with mkdocs tool
*/
type markdownTemplate struct {
	provider analyserTypes.IProvider
	outDir string
}

func (d markdownTemplate) Serve(port string) error {
	panic("implement me")
}

func (d *markdownTemplate) Build(outDir string) error {
	d.outDir = outDir
	for _, pkg := range d.provider.GetPackages() {
		if err := d.processPackage(pkg); err != nil {
			return err
		}
	}

	return nil
}

func (d *markdownTemplate) processPackage(pkg *ast.Package) error {
	buf := types.TmplPkg{
		Pkg: pkg.Name,
	}
	outputPath, err := templates.GetOutputPathForPackage(d.provider.GetRootDir(), pkg)
	if err != nil {
		return err
	}
	outputPath = path.Join(d.outDir, outputPath)

	// in the markdown, we want path which starts with separator
	buf.PkgPath = string(os.PathSeparator) + outputPath

	funcs := d.provider.GetObjects(ast.Fun)
	for _, file := range pkg.Files {
		for _, obj := range file.Scope.Objects {
			if _, exists := funcs[obj]; !exists {
				continue
			}
			funcData := types.TmplFunc{
				Name:    obj.Name,
				Comment: funcs[obj].CommentGroup.Text(),
			}
			funcData.BuildSignature(obj.Decl.(*ast.FuncDecl))

			buf.Funcs = append(buf.Funcs, funcData)
		}
	}

	if outputPath != "" {
		if err := os.MkdirAll(outputPath, os.ModePerm); err != nil {
			return err
		}
	}
	outFile, err := os.Create(path.Join(outputPath, buf.Pkg) + ".md")
	if err != nil {
		return err
	}
	return mdTemplate.Execute(outFile, buf)
}

func NewTemplate(provider analyserTypes.IProvider) types.ITemplater {
	return &markdownTemplate{
		provider,
		"",
	}
}

func init() {
	mdTemplate = template.Must(template.New("mdfile").Parse(tmpl))
}
