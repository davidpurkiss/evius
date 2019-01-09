package gogen

import (
	"evius/util/directory"
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

	astPackage := pkgs[name]

	files := make([]*File, 0)

	pkg := &Package{name, packagePath, files, astPackage, fset}

	if astPackage != nil {

		for path, f := range astPackage.Files {
			files = append(files, NewFile(path, pkg, f))
		}
	}

	pkg.files = files

	return pkg, nil
}

// GetFilePath gets the full file path for a given file name in a package
func (pkg *Package) GetFilePath(name string) string {

	absolutePackagePath, _ := filepath.Abs(pkg.path)

	strs := []string{}
	strs = append(strs, name)
	strs = append(strs, ".go")

	return path.Join(absolutePackagePath, strings.Join(strs, ""))
}

// OpenFile opens a new file in the package
func (pkg *Package) OpenFile(name string) (*File, error) {
	astFile := pkg._package.Files[name]
	filePath := pkg.GetFilePath(name)

	file := NewFile(filePath, pkg, astFile)
	pkg.files = append(pkg.files, file)

	return file, nil
}

// CreateFile creates a new file
func (pkg *Package) CreateFile(name string) (*File, error) {

	filePath := pkg.GetFilePath(name)

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

		for _, file := range pkg.files {
			if file.name == name {
				newFile = file
			}
		}
	} else {
		pkg._package.Files[name] = newAstFile
		newFile = &File{name, filePath, make([]*Type, 0), make([]*Struct, 0), make([]*Interface, 0), make([]*Func, 0), pkg, newAstFile}
		pkg.files = append(pkg.files, newFile)
	}

	return newFile, nil
}

// GetFile retrieves an existing and open file from the package
func (pkg *Package) GetFile(name string) *File {
	for _, file := range pkg.files {
		if file.name == name {
			return file
		}
	}

	return nil
}

// RemoveFile removes an existing file
func (pkg *Package) RemoveFile(name string) error {
	filePath := pkg.GetFilePath(name)

	if !directory.Exists(filePath) {
		return fmt.Errorf("The file '%s' does not exist", name)
	}

	for index, file := range pkg.files {
		if file.name == name {
			pkg.files = append(pkg.files[:index], pkg.files[index+1:]...)
			break
		}
	}

	return directory.Remove(filePath)
}

// RenameFile renames an existing file
func (pkg *Package) RenameFile(oldName string, newName string) (*File, error) {

	filePath := pkg.GetFilePath(oldName)
	newFilePath := pkg.GetFilePath(newName)

	if !directory.Exists(filePath) {
		return nil, fmt.Errorf("The file '%s' does not exist", oldName)
	}

	file := pkg.GetFile(oldName)

	if err := directory.Rename(filePath, newFilePath); err != nil {
		return nil, err
	}

	file.name = newName
	file.path = newFilePath

	return file, nil
}
