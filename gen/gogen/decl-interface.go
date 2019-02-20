package gogen

import "reflect"

// Interface represents a go interface and defines its attributes
type Interface struct {
	name        string
	description string
}

// Name returns the name of the interface
func (intfc Interface) Name() string {
	return intfc.name
}

// Description returns the description of the interface
func (intfc Interface) Description() string {
	return intfc.description
}

// Type returns the type name
func (intfc Interface) Type() string {
	return reflect.TypeOf(intfc).Name()
}
