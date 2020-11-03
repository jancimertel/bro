package analyser

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

const testingDir = "../examples/teststructure/"

func TestAnalyser_Process(t *testing.T) {
	a := NewProjectAnalyser(testingDir)

	t.Run("text filter", func(t *testing.T) {
		a.Process()
		assert.NotEmpty(t, a.codeParser.GetObjects(ast.Fun))
	})
}