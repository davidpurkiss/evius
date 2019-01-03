package gogen

import (
	"evius/util/directory"
	"fmt"
	"path"
)

// Workspace defines the attributes of a go workspace
type Workspace struct {
	name string
	path string
}

// CreateWorkspace creates a new go workspace
func (workspace *Workspace) CreateWorkspace() (*Workspace, error) {

	return workspace, directory.Create(workspace.path)
}

// CreateRepository creates a new repository in the go workspace
func (workspace Workspace) CreateRepository(name string) (*Repository, error) {

	repoPath := path.Join(workspace.path, name)
	if directory.Exists(repoPath) {
		return &Repository{}, fmt.Errorf("The repo '%s' already exists", name)
	}
	repo := &Repository{name: name, path: repoPath}
	return repo, directory.Create(repo.path)
}

// OpenRepository opens an existing repository
func (workspace *Workspace) OpenRepository(name string) (*Repository, error) {

	repoPath := path.Join(workspace.path, name)
	if !directory.Exists(repoPath) {
		return nil, fmt.Errorf("The repo '%s' does not exist", name)
	}

	return &Repository{name: name, path: repoPath}, nil
}

// RemoveRepository removes an existing repository from the go workspace
func (workspace *Workspace) RemoveRepository(name string) error {

	repoPath := path.Join(workspace.path, name)
	if !directory.Exists(repoPath) {
		return fmt.Errorf("The repo '%s' does not exist", name)
	}
	return directory.Remove(repoPath)
}

// RenameRepository renames an existing repository in a go workspace
func (workspace Workspace) RenameRepository(oldName string, newName string) (Repository, error) {
	repoPath := path.Join(workspace.path, oldName)
	repo := Repository{name: oldName, path: repoPath}
	newRepoPath := path.Join(workspace.path, newName)

	if directory.Exists(newRepoPath) {
		return repo, fmt.Errorf("The repo '%s' already exists", newName)
	}
	err := directory.Rename(repoPath, newRepoPath)

	if err == nil {
		repo.name = newName
		repo.path = newRepoPath
	}

	return repo, err
}
