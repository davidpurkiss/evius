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
	_typeDecls map[string]*ast.GenDecl
}

// NewFile initializes a new file from a ast.File instance
func NewFile(filePath string, pkg *Package, astFile *ast.File) *File {

	_, name := path.Split(filePath)
	name = strings.Replace(name, path.Ext(name), "", 1)

	return &File{name, filePath, make([]*Type, 0), make([]*Struct, 0), make([]*Interface, 0), make([]*Func, 0), pkg, astFile, make(map[string]*ast.GenDecl, 0)}
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

// GetType retrieves an existing type from the file using its name
func (file *File) GetType(name string) *Type {
	for _, typ := range file.types {
		if typ.name == name {
			return typ
		}
	}

	return nil
}

// AddType adds a new type to the file
func (file *File) AddType(name string, description string, baseTypeName string) (*Type, error) {

	if existingType := file.GetType(name); existingType != nil {
		return nil, fmt.Errorf("The type '%s' already exists", name)
	}

	newType := NewType(name, description, baseTypeName)
	file.types = append(file.types, newType)

	decl := &ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{newType._type}}
	file._file.Decls = append(file._file.Decls, decl)
	file._typeDecls[name] = decl

	file.Save()

	return newType, nil
}

// RenameType renames an existing type using its old name
func (file *File) RenameType(oldName string, newName string) {
	typ := file.GetType(oldName)

	typ.name = newName
	typ.SetName(newName)

	file.Save()
}

// RemoveType renames an existing type using its old name
func (file *File) RemoveType(name string) error {
	if existingType := file.GetType(name); existingType == nil {
		return fmt.Errorf("The type '%s' does not exist", name)
	}

	decl := file._typeDecls[name]

	if decl == nil {
		return fmt.Errorf("No declaration was found for type '%s'", name)
	}

	// Remove the Delcaration from the ast
	for i, d := range file._file.Decls {
		if d == decl {
			file._file.Decls = append(file._file.Decls[:i], file._file.Decls[i+1:]...)
		}
	}

	// Remove the file from the file types
	for i, t := range file.types {
		if t.name == name {
			file.types = append(file.types[:i], file.types[i+1:]...)
		}
	}

	return nil
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
