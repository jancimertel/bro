package text

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

type dummyProvider struct {
	
}

func (d dummyProvider) GetObjects(kind ast.ObjKind) map[*ast.Object]*analyserTypes.Object {
	return nil
}

func (d dummyProvider) GetPackages() map[string]*ast.Package {
	return nil
}

func (d dummyProvider) Dump() *analyserTypes.FullDump {
	return nil
}

func Test_textTemplate_Build(t *testing.T) {
	template := NewTemplate(&dummyProvider{})
	t.Run("basic build should not fail", func(t *testing.T) {
		assert.Nil(t, template.Build())
	})
}