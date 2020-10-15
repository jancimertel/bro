package analyser

import (
	"bitbucket.org/jmertel/bro/analyser/comments"
	"bitbucket.org/jmertel/bro/analyser/structure"
	"bitbucket.org/jmertel/bro/analyser/types"
	"go/ast"
)

type Analyser struct {
	codeParser     types.ICodeParser
	commentsParser types.ICommentsExtractor
}

func (a *Analyser) GetPackages() map[string]*ast.Package {
	return a.codeParser.GetPackages()
}

func (a *Analyser) GetObjects(kind ast.ObjKind) map[*ast.Object]*types.Object {
	return a.codeParser.GetObjects(kind)
}

func NewProjectAnalyser(pathToProject string) Analyser {
	return Analyser{
		codeParser:     structure.NewParser(pathToProject),
		commentsParser: comments.NewComments(),
	}
}

func (a *Analyser) Process() {
	a.codeParser.ParseProject(nil)
	a.codeParser.MakeReversedRefs()
}
