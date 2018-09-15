package gogen

// Repository defines the attributes of a go repository
type Repository struct {
	name string
	path string
}

// CreatePackage creates a new package in the go repository
func (repository Repository) CreatePackage(name string) {

}

// RemovePackage removes an existing package from the go repository
func (repository Repository) RemovePackage(name string) {

}

// RenamePackage renames an existing package in a go repository
func (repository Repository) RenamePackage(oldName string, newName string) {

}
