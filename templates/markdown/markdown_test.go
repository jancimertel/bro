package markdown

import (
	"bitbucket.org/jmertel/bro/analyser"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Test_textTemplate_processPackage(t *testing.T) {
	provider := analyser.NewProjectAnalyser("../../examples/teststructure")
	provider.Process()
	template := NewTemplate(&provider).(*markdownTemplate)
	dummyPackages := template.provider.GetPackages()
	rootPackage := dummyPackages["rootpkg"]

	t.Run("processPackage", func(t *testing.T) {
		assert.NotPanics(t, func() {
			assert.Nil(t, template.processPackage(rootPackage))
		})
		assert.True(t, fileExists(path.Join(os.TempDir(), rootPackage.Name, rootPackage.Name + ".md")))
	})
}