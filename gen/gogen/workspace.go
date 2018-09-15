package gogen

// Workspace defines the attributes of a go workspace
type Workspace struct {
	name string
	path string
}

// CreateRepository creates a new repository in the go workspace
func (workspace Workspace) CreateRepository(name string) {

}

// RemoveRepository removes an existing repository from the go workspace
func (workspace Workspace) RemoveRepository(name string) {

}

// RenameRepository renames an existing repository in a go workspace
func (workspace Workspace) RenameRepository(oldName string, newName string) {

}
