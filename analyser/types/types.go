package types

import (
	"go/ast"
	"os"
)

type ICommentsExtractor interface {

}

type CommentNode struct {
	*ast.Scope
}

type ICodeParser interface {
	ParseProject(filter func(info os.FileInfo) bool)
	GetComments(pkg string) []*CommentNode
}

