package templates

import (
	"errors"
	"go/ast"
	"os"
	"strings"
)

func GetOutputPathForPackage(rootPath string, pkg *ast.Package) (string, error) {
	if len(pkg.Files) == 0 {
		return "", errors.New("package should not exist - empty file list")
	}

	dirPrefix := strings.TrimPrefix(rootPath, "./")

	var realPath string
	for path := range pkg.Files {
		realPath = path
		break
	}

	pkgPath := strings.TrimPrefix(realPath, dirPrefix)
	if strings.Contains(pkgPath, string(os.PathSeparator)) {
		pkgPath = pkgPath[0:strings.LastIndex(pkgPath, string(os.PathSeparator))]
	} else {
		pkgPath = ""
	}

	return pkgPath, nil
}