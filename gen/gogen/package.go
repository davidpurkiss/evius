package gogen

import (
	"evius/util"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Package defines the attributes of a go package
type Package struct {
	name     string
	path     string
	files    []*File
	_package *ast.Package
	_fset    *token.FileSet
}

// OpenPackage initializes a new package structure from a directory
func OpenPackage(packagePath string) (*Package, error) {

	absolutePath, _ := filepath.Abs(packagePath)
	_, name := path.Split(absolutePath)

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, absolutePath, nil, parser.ParseComments)

	if err != nil {
		return nil, err
	}

	pkg := pkgs[name]

	files := make([]*File, 0)

	for path, f := range pkg.Files {
		files = append(files, OpenFile(path, f))
	}

	return &Package{name, packagePath, files, pkg, fset}, nil
}

// CreateFile creates a new file
func (pkg Package) CreateFile(name string) (*File, error) {
	filePath := path.Join(pkg.path, name)
	if directory.Exists(filePath) {
		return &File{}, fmt.Errorf("The file '%s' already exists", name)
	}

	newFile := &ast.File{&ast.CommentGroup{}, pkg._package.Pos(), ast.NewIdent(name), []ast.Decl{}, ast.NewScope(pkg._package.Scope), []*ast.ImportSpec{}, nil, []*ast.CommentGroup{}}

	pkg._package.Files[name] = newFile

	strs := []string{}
	strs = append(strs, name)
	strs = append(strs, ".go")

	f, _ := os.Create(strings.Join(strs, ""))
	defer f.Close()

	if err := printer.Fprint(f, pkg._fset, newFile); err != nil {
		log.Fatal(err)
	}

	return &File{name, filePath, newFile}, nil
}

// // RemoveItem removes an existing item
// func RemoveItem(name string) error {

// }

// // RenameItem renames an existing item
// func RenameItem(oldName string, newName string) (Item, error) {

// }
