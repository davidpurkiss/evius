package gogen

import (
	"go/ast"
)

// File defines the attributes of a go file (or code item)
type File struct {
	name       string
	path       string
	types      []*Type
	structs    []*Struct
	interfaces []*Interface
	functions  []*Func
	_file      *ast.File
}

// OpenFile initializes a new file from a ast.File instance
func OpenFile(path string, astFile *ast.File) *File {

	return &File{astFile.Name.String(), path, make([]*Type, 0), make([]*Struct, 0), make([]*Interface, 0), make([]*Func, 0), astFile}
}

// Items returns a list of all root level items in a file. Types, Structs, Interfaces and Functions
func (file File) Items() []Item {

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
