package gogen

import (
	"go/ast"
	"reflect"
)

// Func represents a go func and defines its attributes
type Func struct {
	name        string
	description string
	receiver    *Field
	params      []*Field
	results     []*Field
	file        *File
	_decl       *ast.FuncDecl
	_body       *ast.BlockStmt
}

// Name returns the name of the function
func (fnc Func) Name() string {
	return fnc.name
}

// Description returns the description of the function
func (fnc Func) Description() string {
	return fnc.description
}

// Type returns the type name
func (fnc Func) Type() string {
	return reflect.TypeOf(fnc).Name()
}
