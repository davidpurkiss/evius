package gogen

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

// Type represents a custom go type and defines its attributes
type Type struct {
	name        string
	description string
	file        *File
	_type       *ast.TypeSpec
	_decl       *ast.GenDecl
}

// NewType creates a new Type ready for population
func NewType(name string, description string, baseTypeName string, file *File) *Type {

	var cg *ast.CommentGroup
	if description != "" {
		cg = &ast.CommentGroup{
			List: []*ast.Comment{&ast.Comment{Text: fmt.Sprint("// ", description), Slash: 0}},
		}
	}

	typeSpec := &ast.TypeSpec{Name: ast.NewIdent(name), Type: ast.NewIdent(baseTypeName)}

	decl := &ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{typeSpec}, Doc: cg}

	return &Type{name: name, description: description, file: file, _type: typeSpec, _decl: decl}
}

// SetName sets or changes the name of a Type
func (tpe *Type) SetName(name string) {
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

	var cg *ast.CommentGroup

	if description != "" {
		cg = &ast.CommentGroup{
			List: []*ast.Comment{&ast.Comment{Slash: 0, Text: fmt.Sprint("// ", description)}},
		}
	}

	tpe._decl.Doc = cg
	tpe.file.Save()
}

// Description returns the description of the type
func (tpe *Type) Description() string {
	return tpe.description
}

// SetType sets or changes the type name of the type
func (tpe *Type) SetType(baseTypeName string) {
	tpe._type.Type = ast.NewIdent(baseTypeName)
	tpe.file.Save()
}

// Type returns the type name
func (tpe *Type) Type() string {
	return reflect.TypeOf(tpe).Name()
}
