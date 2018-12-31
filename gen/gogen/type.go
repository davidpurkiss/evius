package gogen

import "reflect"

// Type represents a custom go type and defines its attributes
type Type struct {
	name        string
	description string
}

// Name returns the name of the type
func (tpe Type) Name() string {
	return tpe.name
}

// Description returns the description of the type
func (tpe Type) Description() string {
	return tpe.description
}

// Type returns the type name
func (tpe Type) Type() string {
	return reflect.TypeOf(tpe).Name()
}
