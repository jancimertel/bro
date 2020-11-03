package analyser

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"go/ast"
)

type DummyProvider struct {

}

func (d DummyProvider) GetObjects(kind ast.ObjKind) map[*ast.Object]*analyserTypes.Object {
	return nil
}

func (d DummyProvider) GetPackages() map[string]*ast.Package {
	return nil
}

func (d DummyProvider) Dump() *analyserTypes.FullDump {
	return nil
}

