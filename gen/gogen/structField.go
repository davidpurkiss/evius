package gogen

import (
	"go/ast"
	"reflect"
)

// StructField represents a field within a go struct
type StructField struct {
	name        string
	description string
	typeName    string
	_field      *ast.Field
	file        *File
}

// NewStructField creates a new StructField
func NewStructField(name string, typeName string, description string, file *File) *StructField {
	astField := &ast.Field{Names: []*ast.Ident{ast.NewIdent(name)}, Type: ast.NewIdent(typeName), Comment: getCommentGroup(description)}
	field := &StructField{name: name, typeName: typeName, description: description, file: file, _field: astField}

	return field
}

// SetName sets or changes the name of a StructField
func (field *StructField) SetName(name string) {
	field.name = name
	field._field.Names = []*ast.Ident{ast.NewIdent(name)}
	field.file.Save()
}

// Name returns the name of the StructField
func (field *StructField) Name() string {
	return field.name
}

// Description returns the description of the StructField
func (field *StructField) Description() string {
	return field.description
}

// SetDescription sets or changes the description of the StructField
func (field *StructField) SetDescription(description string) {
	field.description = description

	field._field.Comment = getCommentGroup(description)
	field.file.Save()
}

// SetType sets or changes the type name of the StructField
func (field *StructField) SetType(baseTypeName string) {
	field._field.Type = ast.NewIdent(baseTypeName)
	field.file.Save()
}

// Type returns the type name
func (field *StructField) Type() string {
	return reflect.TypeOf(field).Name()
}
