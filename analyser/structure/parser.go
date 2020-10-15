package structure

import (
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
	packageFiles    map[string]*ast.Package
	rootPath        string
	reversedObjects map[ast.ObjKind]map[*ast.Object]*types.Object
}

func NewParser(dir string) *broStructure {
	revObjects := make(map[ast.ObjKind]map[*ast.Object]*types.Object)
	for kind := ast.Bad; kind <= ast.Lbl; kind++ {
		revObjects[kind] = make(map[*ast.Object]*types.Object)
	}

	return &broStructure{
		make(map[string]*ast.Package),
		dir,
		revObjects,
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

// using the go parser, this method will scan all dirs
func (p *broStructure) ParseProject(filter func(info os.FileInfo) bool) {
	for _, dir := range p.listDirs() {
		fileset := token.NewFileSet()
		mapped, _ := parser.ParseDir(fileset, dir, filter, parser.AllErrors|parser.ParseComments)
		for key, val := range mapped {
			p.packageFiles[key] = val
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

func (p *broStructure) GetObjects(kind ast.ObjKind) map[*ast.Object]*types.Object {
	return p.reversedObjects[kind]
}

// objects from ast package are not enough - we want to store additional information to them.
// This is one of the possibilities to do so
func (p *broStructure) MakeReversedRefs() {
	for _, pkg := range p.packageFiles {
		for _, file := range pkg.Files {
			for _, obj := range file.Scope.Objects {
				p.reversedObjects[obj.Kind][obj] = &types.Object{
					Object:       obj,
					CommentGroup: findCommentForObject(obj, file.Comments),
				}
			}
		}
	}

	return
}

// comments are not part of Scope.Objects, they are texts in with position
// according to the position the comment can be assigned to specific comment
func findCommentForObject(object *ast.Object, commentgrps []*ast.CommentGroup) *ast.CommentGroup {
	for i := range commentgrps {
		// iterating bottom-top so we can find the first comment that is above the object
		commentgrp := commentgrps[len(commentgrps) - 1 - i]

		// general condition - we are trying to find the first comment just above the object
		if object.Pos() > commentgrp.End() {
			switch object.Kind {
			case ast.Fun:
				// "<end><newline>func <func name>"
				if object.Pos()-6 > commentgrp.End() {
					return nil
				}
			default:
				fmt.Println(fmt.Sprintf("findObjectForComment: type not implemented: %v", object.Kind))
				return nil
			}
			return commentgrp
		}
	}

	return nil
}
