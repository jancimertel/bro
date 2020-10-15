package _default

import (
	analyserTypes "bitbucket.org/jmertel/bro/analyser/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_defaultTemplate_Build(t *testing.T) {
	template := defaultTemplate{
		provider: analyserTypes.IProvider(struct{}{}),
	}
	t.Run("basic build should not fail", func(t *testing.T) {
		assert.Nil(t, template.Build())
	})
}