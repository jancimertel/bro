package analyser

import (
	"bitbucket.org/jmertel/bro/analyser/comments"
	"bitbucket.org/jmertel/bro/analyser/structure"
	"bitbucket.org/jmertel/bro/analyser/types"
)

type Analyser struct {
	codeParser     types.ICodeParser
	commentsParser types.ICommentsExtractor
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
