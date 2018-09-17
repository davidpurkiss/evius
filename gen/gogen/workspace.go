package gogen

import (
	"os"
	"path"
)

// Workspace defines the attributes of a go workspace
type Workspace struct {
	name string
	path string
}

// CreateWorkspace creates a new go workspace
func (workspace Workspace) CreateWorkspace() (Workspace, error) {

	return workspace, os.MkdirAll(workspace.path, os.ModePerm)
}

// CreateRepository creates a new repository in the go workspace
func (workspace Workspace) CreateRepository(name string) (Repository, error) {

	repo := Repository{name: name, path: path.Join(workspace.path, name)}
	return repo, os.MkdirAll(repo.path, os.ModePerm)
}

// RemoveRepository removes an existing repository from the go workspace
func (workspace Workspace) RemoveRepository(name string) error {

	repoPath := path.Join(workspace.path, name)
	return os.Remove(repoPath)
}

// RenameRepository renames an existing repository in a go workspace
func (workspace Workspace) RenameRepository(oldName string, newName string) {

}
