package types

import (
	"fmt"
	"go/ast"
	"strings"
)

type ITemplater interface {
	Serve(port string) error
	Build() error
}

type TmplPkg struct {
	Pkg        string
	PkgPath    string
	PkgComment string
	Funcs      []TmplFunc
}

type TmplFunc struct {
	Name      string
	Comment   string
	Signature TmplFuncSignature
}

func (ft *TmplFunc) BuildSignature(params *ast.FuncDecl) {
	ft.Signature = TmplFuncSignature{
		Text:   "",
		Params: nil,
	}

	for _, d := range params.Type.Params.List {
		param := TmplParam{
			Type:  fmt.Sprintf("%s", d.Type),
		}

		for _, ident := range d.Names {
			param.Names = append(param.Names, ident.Name)
		}
		ft.Signature.Params = append(ft.Signature.Params, param)
	}

	for i, param := range ft.Signature.Params {
		if i > 0 {
			ft.Signature.Text += ", "
		}
		ft.Signature.Text += fmt.Sprintf("%s %s", strings.Join(param.Names, ", "), param.Type)
	}
}

type TmplFuncSignature struct {
	Text string
	Params []TmplParam
}

type TmplParam struct {
	Names []string
	Type string
}
