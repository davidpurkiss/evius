package gogen

// Item is an interface used to describe the various code generation structures such as, Type, Struct, Interface and Func
type Item interface {
	Name() string
	Description() string
	Type() string
}
