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
}

type IProvider interface {

}