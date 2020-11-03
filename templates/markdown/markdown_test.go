package markdown

import (
	"bitbucket.org/jmertel/bro/analyser"
	"bitbucket.org/jmertel/bro/templates"
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
	if info == nil {
		return false
	}
	return !info.IsDir()
}

func Test_textTemplate_processPackage(t *testing.T) {
	rootPath := "../../examples/teststructure"
	provider := analyser.NewProjectAnalyser(rootPath)
	provider.Process()
	template := NewTemplate(&provider).(*markdownTemplate)
	dummyPackages := template.provider.GetPackages()
	rootPackage := dummyPackages["rootpkg"]

	t.Run("processPackage", func(t *testing.T) {
		assert.NotPanics(t, func() {
			assert.Nil(t, template.processPackage(rootPackage))
		})
		outputPath, err := templates.GetPathForPackage(rootPath, rootPackage)
		assert.Nil(t, err)
		pathToMdFile := path.Join(outputPath, rootPackage.Name) + ".md"
		assert.True(t, fileExists(pathToMdFile))
		os.RemoveAll(pathToMdFile)
	})
}