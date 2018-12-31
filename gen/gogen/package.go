package gogen

import (
	"evius/util"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
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

	if pkg != nil {

		for path, f := range pkg.Files {
			files = append(files, OpenFile(path, f))
		}
	}

	return &Package{name, packagePath, files, pkg, fset}, nil
}

// CreateFile creates a new file
func (pkg Package) CreateFile(name string) (*File, error) {

	absolutePackagePath, _ := filepath.Abs(pkg.path)

	strs := []string{}
	strs = append(strs, name)
	strs = append(strs, ".go")

	filePath := path.Join(absolutePackagePath, strings.Join(strs, ""))

	if directory.Exists(filePath) {
		return &File{}, fmt.Errorf("The file '%s' already exists", name)
	}

	newAstFile := &ast.File{Doc: &ast.CommentGroup{}, Package: 0, Name: ast.NewIdent(pkg.name), Decls: []ast.Decl{}, Scope: nil, Imports: []*ast.ImportSpec{}, Comments: []*ast.CommentGroup{}}

	f, _ := os.Create(filePath)
	if err := printer.Fprint(f, pkg._fset, newAstFile); err != nil {
		return nil, err
	}
	f.Close()

	var newFile *File

	if pkg._package == nil {
		nPkg, err := OpenPackage(pkg.path)
		if err != nil {
			return nil, err
		}
		pkg._package = nPkg._package
		pkg._fset = nPkg._fset
		pkg.files = nPkg.files

		for _, f := range pkg.files {
			if f.name == name {
				newFile = f
			}
		}
	} else {
		pkg._package.Files[name] = newAstFile
		newFile = &File{name, filePath, make([]*Type, 0), make([]*Struct, 0), make([]*Interface, 0), make([]*Func, 0), newAstFile}
		pkg.files = append(pkg.files, newFile)
	}

	return newFile, nil
}

// // RemoveItem removes an existing item
// func RemoveItem(name string) error {

// }

// // RenameItem renames an existing item
// func RenameItem(oldName string, newName string) (Item, error) {

// }
