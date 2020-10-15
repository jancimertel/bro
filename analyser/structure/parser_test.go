package structure

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

const testingDir = "../../examples/teststructure/"

func goFilesFilter(filename string) (accept bool) {
	return strings.HasSuffix(filename, ".go")
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if pathParts := strings.Split(path, string(os.PathSeparator)); goFilesFilter(pathParts[len(pathParts)-1]) {
			*files = append(*files, path)
		}
		return nil
	}
}

func getAllGoFiles(dir string) (files []string) {
	if err := filepath.Walk(dir, visit(&files)); err != nil {
		panic(err)
	}
	for i, path := range files {
		files[i] = strings.Replace(path, dir, "", 1)
	}
	sort.Strings(files)
	return
}

func Test_broStructure_parseProject(t *testing.T) {

	p := NewParser(testingDir)

	t.Run("parse project non empty/nonpanic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			p.ParseProject(nil)
		})
		assert.NotEmpty(t, p.packageFiles)
	})

	t.Run("equal to file list", func(t *testing.T) {
		expected := getAllGoFiles(testingDir)
		p.ParseProject(nil)
		assert.Equal(t, expected, p.listFiles())
	})
}

func Test_broStructure_listDirs(t *testing.T) {
	// as the implementation is rather low level, i hardcoded those results - removed prepended path parts to be sure
	expected := []string{"examples/teststructure", "examples/teststructure/pkg1", "examples/teststructure/pkg2"}
	var out []string

	t.Run("listDirs hardcoded", func(t *testing.T) {
		p := NewParser(testingDir)

		out = p.listDirs()
		for i, fullPath := range out {
			pathParts := strings.Split(fullPath, string(os.PathSeparator))
			for j, part := range pathParts {
				if part == "examples" {
					out[i] = strings.Join(pathParts[j:], string(os.PathSeparator))
					if pathParts[len(pathParts) - 1] == "" {
						out[i] = out[i][:len(out[i]) - 1]
					}
					break
				}
				if j == len(pathParts) - 1 {
					panic(fmt.Sprintf("full path cannot be trimmed: %s", fullPath))
				}
			}
		}
		assert.Equal(t, expected, out)
	})
}
