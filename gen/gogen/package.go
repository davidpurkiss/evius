package gogen

// Package defines the attributes of a go package
type Package struct {
	name string
	path string
	//_package ast.Package
}

func NewPackage(name string, path string) *Package {
	packg := Package{name, path}
	//packg._package = ast.NewPackage()
	return &packg
}

// // CreateItem creates a new item
// func CreateItem(name string) (Item, error) {

// }

// // RemoveItem removes an existing item
// func RemoveItem(name string) error {

// }

// // RenameItem renames an existing item
// func RenameItem(oldName string, newName string) (Item, error) {

// }
