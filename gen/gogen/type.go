package gogen

import (
	"fmt"
	"go/ast"
	"reflect"
)

// Type represents a custom go type and defines its attributes
type Type struct {
	name        string
	description string
	_type       *ast.TypeSpec
}

// NewType creates a new Type ready for population
func NewType(name string, description string, baseTypeName string) *Type {

	var cg *ast.CommentGroup
	if description != "" {
		cg = &ast.CommentGroup{
			List: []*ast.Comment{&ast.Comment{Text: fmt.Sprint("// ", description), Slash: 0}},
		}
	}

	typeSpec := &ast.TypeSpec{Name: ast.NewIdent(name), Type: ast.NewIdent(baseTypeName), Doc: cg}

	return &Type{name: name, description: description, _type: typeSpec}
}

// SetName sets or changes the name of a Type
func (tpe *Type) SetName(name string) {
	tpe._type.Name = ast.NewIdent(name)
}

// Name returns the name of the type
func (tpe *Type) Name() string {
	return tpe.name
}

// SetDescription sets or changes the description of the type
func (tpe *Type) SetDescription(description string) {
	var cg *ast.CommentGroup

	if description != "" {
		cg = &ast.CommentGroup{
			List: []*ast.Comment{&ast.Comment{Slash: 0, Text: description}},
		}
	}

	tpe._type.Doc = cg
}

// Description returns the description of the type
func (tpe *Type) Description() string {
	return tpe.description
}

// SetType sets or changes the type name of the type
func (tpe *Type) SetType(baseTypeName string) {
	tpe._type.Type = ast.NewIdent(baseTypeName)
}

// Type returns the type name
func (tpe *Type) Type() string {
	return reflect.TypeOf(tpe).Name()
}
