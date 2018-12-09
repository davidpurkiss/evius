package gogen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"path/filepath"
)

// Package defines the attributes of a go package
type Package struct {
	name     string
	path     string
	_package *ast.Package
}

// NewPackage initializes a new package structure from a directory
func NewPackage(packagePath string) (*Package, error) {

	absolutePath, _ := filepath.Abs(packagePath)
	_, name := path.Split(absolutePath)

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, absolutePath, nil, parser.ParseComments)

	if err != nil {
		return nil, err
	}

	return &Package{name, packagePath, pkgs[name]}, nil
}

// // CreateItem creates a new item
// func CreateItem(name string) (Item, error) {

// }

// // RemoveItem removes an existing item
// func RemoveItem(name string) error {

// }

// // RenameItem renames an existing item
// func RenameItem(oldName string, newName string) (Item, error) {

// }
