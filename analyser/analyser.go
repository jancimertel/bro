package analyser

import (
	"bitbucket.org/jmertel/bro/analyser/comments"
	"bitbucket.org/jmertel/bro/analyser/structure"
)

type Analyser struct {
	codeParser     ICodeParser
	commentsParser ICommentsExtractor
}

func NewProjectAnalyser(pathToProject string) Analyser {
	return Analyser{
		codeParser:     structure.NewParser(pathToProject),
		commentsParser: comments.NewComments(),
	}
}

func (a *Analyser) Process() {
	a.codeParser.ParseProject(nil)
}
