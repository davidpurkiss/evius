package gogen

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	"path"
	"strings"
)

// File defines the attributes of a go file (or code item)
type File struct {
	name       string
	path       string
	types      []*Type
	structs    []*Struct
	interfaces []*Interface
	functions  []*Func
	pkg        *Package
	_file      *ast.File
}

// OpenFile initializes a new file from a ast.File instance
func OpenFile(filePath string, pkg *Package, astFile *ast.File) *File {

	_, name := path.Split(filePath)
	name = strings.Replace(name, path.Ext(name), "", 1)

	return &File{name, filePath, make([]*Type, 0), make([]*Struct, 0), make([]*Interface, 0), make([]*Func, 0), pkg, astFile}
}

// Save writes the current ast to disk
func (file *File) Save() error {

	f, _ := os.OpenFile(file.path, os.O_RDWR, 0755)
	defer f.Close()
	if err := printer.Fprint(f, file.pkg._fset, file._file); err != nil {
		return err
	}

	return nil
}

// AddType adds a new type to the file
func (file *File) AddType(name string, description string, baseTypeName string) (*Type, error) {
	newType := NewType(name, description, baseTypeName)
	file.types = append(file.types, newType)

	decl := &ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{newType._type}}
	file._file.Decls = append(file._file.Decls, decl)

	return newType, nil
}

// AddStruct adds a new struct to the file
func (file *File) AddStruct(name string, description string) (*Struct, error) {
	return nil, fmt.Errorf("Function not implemented")
}

// AddInterface adds a new struct to the file
func (file *File) AddInterface(name string, description string) (*Interface, error) {
	return nil, fmt.Errorf("Function not implemented")
}

// AddFunction adds a new function to the file
func (file *File) AddFunction(name string, description string) (*Func, error) {
	return nil, fmt.Errorf("Function not implemented")
}

// RemoveType removes a type from the file
func (file *File) RemoveType(*Type) error {
	return fmt.Errorf("Function not implemented")
}

// RemoveStruct removes a struct from the file
func (file *File) RemoveStruct(*Struct) error {
	return fmt.Errorf("Function not implemented")
}

// RemoveInterface removes a struct from the file
func (file *File) RemoveInterface(*Interface) error {
	return fmt.Errorf("Function not implemented")
}

// RemoveFunction removes a function from the file
func (file *File) RemoveFunction(*Func) error {
	return fmt.Errorf("Function not implemented")
}

// Items returns a list of all root level items in a file. Types, Structs, Interfaces and Functions
func (file *File) Items() []Item {

	items := make([]Item, 0)

	for _, t := range file.types {
		items = append(items, t)
	}

	for _, s := range file.structs {
		items = append(items, s)
	}

	for _, i := range file.interfaces {
		items = append(items, i)
	}

	for _, f := range file.functions {
		items = append(items, f)
	}

	return items
}
