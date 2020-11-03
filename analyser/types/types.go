package types

import (
	"go/ast"
	"os"
)

type ICommentsExtractor interface {

}

type Object struct {
	*ast.Object
	*ast.CommentGroup
}

type ICodeParser interface {
	ParseProject(filter func(info os.FileInfo) bool)
	MakeReversedRefs()
	GetObjects(kind ast.ObjKind) map[*ast.Object]*Object
	GetPackages() map[string]*ast.Package
	GetRootDir() string
}

type IProvider interface {
	GetObjects(kind ast.ObjKind) map[*ast.Object]*Object
	GetPackages() map[string]*ast.Package
	Dump() *FullDump
	GetRootDir() string
}

type Function struct {
	Name string `json:"name"`
}

type Package struct {
	Name string `json:"name"`
	Functions []Function `json:"functions"`
}

type FullDump struct {
	Packages []Package `json:"objects"`
}
