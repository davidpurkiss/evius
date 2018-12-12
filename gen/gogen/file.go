package gogen

import (
	"go/ast"
)

// File defines the attributes of a go file (or code item)
type File struct {
	name  string
	path  string
	_file *ast.File
}

// OpenFile initializes a new file from a ast.File instance
func OpenFile(path string, astFile *ast.File) *File {

	return &File{astFile.Name.String(), path, astFile}
}
