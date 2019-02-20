package gogen

import (
	"go/ast"
	"go/token"
	"reflect"
)

// Type represents a custom go type and defines its attributes
type Type struct {
	name         string
	description  string
	baseTypeName string
	file         *File
	_type        *ast.TypeSpec
	_decl        *ast.GenDecl
}

// NewType creates a new Type ready for population
func NewType(name string, description string, baseTypeName string, file *File) *Type {

	cg := getCommentGroup(description)

	typeSpec := &ast.TypeSpec{Name: ast.NewIdent(name), Type: ast.NewIdent(baseTypeName)}

	decl := &ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{typeSpec}, Doc: cg}

	return &Type{name: name, description: description, baseTypeName: baseTypeName, file: file, _type: typeSpec, _decl: decl}
}

// SetName sets or changes the name of a Type
func (tpe *Type) SetName(name string) {
	tpe.name = name
	tpe._type.Name = ast.NewIdent(name)

	tpe.file.Save()
}

// Name returns the name of the type
func (tpe *Type) Name() string {
	return tpe.name
}

// SetDescription sets or changes the description of the type
func (tpe *Type) SetDescription(description string) {
	tpe.description = description

	tpe._decl.Doc = getCommentGroup(description)

	tpe.file.Save()
}

// Description returns the description of the type
func (tpe *Type) Description() string {
	return tpe.description
}

// SetType sets or changes the type name of the type
func (tpe *Type) SetType(baseTypeName string) {
	tpe.baseTypeName = baseTypeName
	tpe._type.Type = ast.NewIdent(baseTypeName)

	tpe.file.Save()
}

// Type returns the type name
func (tpe *Type) Type() string {
	return reflect.TypeOf(tpe).Name()
}
