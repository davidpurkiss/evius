package gogen

import (
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
	if description == "" {
		cg = &ast.CommentGroup{
			List: []*ast.Comment{&ast.Comment{Text: description}},
		}
	}

	typeSpec := &ast.TypeSpec{Name: ast.NewIdent(name), Type: ast.NewIdent(baseTypeName), Doc: cg}

	return &Type{name: name, description: description, _type: typeSpec}
}

// Name returns the name of the type
func (tpe *Type) Name() string {
	return tpe.name
}

// Description returns the description of the type
func (tpe *Type) Description() string {
	return tpe.description
}

// Type returns the type name
func (tpe *Type) Type() string {
	return reflect.TypeOf(tpe).Name()
}
