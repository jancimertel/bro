package analyser

import (
	"bitbucket.org/jmertel/bro/analyser/comments"
	"bitbucket.org/jmertel/bro/analyser/structure"
	"bitbucket.org/jmertel/bro/analyser/types"
)

type Analyser struct {
	types.ICodeParser
	types.ICommentsExtractor
}

func NewProjectAnalyser(pathToProject string) Analyser {
	return Analyser{
		ICodeParser:        structure.NewParser(pathToProject),
		ICommentsExtractor: comments.NewComments(),
	}
}

func (a *Analyser) Process() {
	a.ParseProject(nil)
	a.MakeReversedRefs()
}

func (a *Analyser) Dump() (out *types.FullDump) {
	out = &types.FullDump{}
	for _, pkg := range	a.GetPackages() {
		out.Packages = append(out.Packages, types.Package{
			Name: pkg.Name,
			Functions: []types.Function{},
		})
	}

	return
}