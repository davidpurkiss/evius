package gogen

import "reflect"

// Struct represents a go struct and defines its attributes
type Struct struct {
	name        string
	description string
}

// Name returns the name of the struct
func (strct Struct) Name() string {
	return strct.name
}

// Description returns the description of the type
func (strct Struct) Description() string {
	return strct.description
}

// Type returns the type name
func (strct Struct) Type() string {
	return reflect.TypeOf(strct).Name()
}
