package gogen

import (
	"fmt"
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

	repoPath := path.Join(workspace.path, name)
	if directoryExists(repoPath) {
		return Repository{}, fmt.Errorf("The repo '%s' already exists", name)
	}
	repo := Repository{name: name, path: repoPath}
	return repo, os.MkdirAll(repo.path, os.ModePerm)
}

// RemoveRepository removes an existing repository from the go workspace
func (workspace Workspace) RemoveRepository(name string) error {

	repoPath := path.Join(workspace.path, name)
	if directoryExists(repoPath) {
		return fmt.Errorf("The repo '%s' does not exist", name)
	}
	return os.Remove(repoPath)
}

// RenameRepository renames an existing repository in a go workspace
func (workspace Workspace) RenameRepository(oldName string, newName string) {

}

func directoryExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
