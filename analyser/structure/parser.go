package structure

import (
	"bitbucket.org/jmertel/bro/analyser"
	"bitbucket.org/jmertel/bro/analyser/types"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type broStructure struct {
	packageFiles map[string]*ast.Package
	rootPath     string
}

func NewParser(dir string) *broStructure {
	return &broStructure{
		nil,
		dir,
	}
}

func (p *broStructure) listDirs() (paths []string) {
	if err := filepath.Walk(p.rootPath, func(files *[]string) filepath.WalkFunc {
		return func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				*files = append(*files, path)
			}

			return nil
		}
	}(&paths)); err != nil {
		panic(err)
	}

	return
}

func (p *broStructure) ParseProject(filter func(info os.FileInfo) bool) {
	p.packageFiles = make(map[string]*ast.Package)
	for _, dir := range p.listDirs() {
		fileset := token.NewFileSet()
		mapped, _ := parser.ParseDir(fileset, dir, filter, parser.AllErrors | parser.ParseComments)
		for key, val := range mapped {
			p.packageFiles[key] = val
			for _, file := range val.Files {

				for k, sc := range file.Scope.Objects {
					fmt.Println("scp", k, sc.Name, sc.Kind)
				}
				for _, commentgrp := range file.Comments {
					fmt.Println("cmnt", commentgrp.Text())
				}
			}
		}

	}
}

func (p *broStructure) listFiles() (files []string) {
	for _, value := range p.packageFiles {
		for file := range value.Files {
			file = strings.Replace(file, p.rootPath, "", 1)
			files = append(files, file)
		}
	}
	sort.Strings(files)

	return
}

func (p *broStructure) GetComments(pkg string) []*types.CommentNode {
	for _, _ := range p.packageFiles[pkg].Scope.Objects {

	}
	return nil
}