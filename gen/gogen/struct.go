package gogen

import (
	"go/ast"
	"go/token"
	"reflect"
)

// StructField represents a field within a go struct
type StructField struct {
	name        string
	description string
	typeName    string
	_field      *ast.Field
}

// Struct represents a go struct and defines its attributes
type Struct struct {
	name        string
	description string
	fields      []*StructField
	file        *File
	_structType *ast.StructType
	_type       *ast.TypeSpec
	_decl       *ast.GenDecl
}

// NewStruct creates a new Struct ready for population
func NewStruct(name string, description string, file *File) *Struct {

	cg := getCommentGroup(description)

	structType := &ast.StructType{Fields: &ast.FieldList{List: make([]*ast.Field, 0)}}
	typeSpec := &ast.TypeSpec{Name: ast.NewIdent(name), Type: structType}

	decl := &ast.GenDecl{Tok: token.TYPE, Specs: []ast.Spec{typeSpec}, Doc: cg}

	return &Struct{name: name, description: description, fields: make([]*StructField, 0), file: file, _structType: structType, _type: typeSpec, _decl: decl}
}

// SetName sets or changes the name of a Struct
func (strct *Struct) SetName(name string) {
	strct._type.Name = ast.NewIdent(name)
	strct.file.Save()
}

// Name returns the name of the Struct
func (strct *Struct) Name() string {
	return strct.name
}

// Description returns the description of the type
func (strct *Struct) Description() string {
	return strct.description
}

// SetDescription sets or changes the description of the Struct
func (strct *Struct) SetDescription(description string) {
	strct.description = description

	strct._decl.Doc = getCommentGroup(description)
	strct.file.Save()
}

// SetType sets or changes the type name of the Struct
func (strct *Struct) SetType(baseTypeName string) {
	strct._type.Type = ast.NewIdent(baseTypeName)
	strct.file.Save()
}

// Type returns the type name
func (strct Struct) Type() string {
	return reflect.TypeOf(strct).Name()
}
