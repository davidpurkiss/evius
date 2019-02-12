package gogen

import (
	"go/ast"
	"reflect"
)

// Field represents a field within a go struct
type Field struct {
	name        string
	description string
	typeName    string
	_field      *ast.Field
	file        *File
}

// NewField creates a new Field
func NewField(name string, typeName string, description string, file *File) *Field {
	astField := &ast.Field{Names: []*ast.Ident{ast.NewIdent(name)}, Type: ast.NewIdent(typeName), Comment: getCommentGroup(description)}
	field := &Field{name: name, typeName: typeName, description: description, file: file, _field: astField}

	return field
}

// SetName sets or changes the name of a Field
func (field *Field) SetName(name string) {
	field.name = name
	field._field.Names = []*ast.Ident{ast.NewIdent(name)}
	field.file.Save()
}

// Name returns the name of the Field
func (field *Field) Name() string {
	return field.name
}

// Description returns the description of the Field
func (field *Field) Description() string {
	return field.description
}

// SetDescription sets or changes the description of the Field
func (field *Field) SetDescription(description string) {
	field.description = description

	field._field.Comment = getCommentGroup(description)
	field.file.Save()
}

// SetType sets or changes the type name of the Field
func (field *Field) SetType(baseTypeName string) {
	field._field.Type = ast.NewIdent(baseTypeName)
	field.file.Save()
}

// Type returns the type name
func (field *Field) Type() string {
	return reflect.TypeOf(field).Name()
}
