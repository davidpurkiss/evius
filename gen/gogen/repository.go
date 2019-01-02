package gogen

import (
	"evius/util/directory"
	"fmt"
	"path"
)

// Repository defines the attributes of a go repository
type Repository struct {
	name string
	path string
}

// CreatePackage creates a new package in a go repository
func (repository Repository) CreatePackage(name string) (*Package, error) {
	packagePath := path.Join(repository.path, name)
	if directory.Exists(packagePath) {
		return &Package{}, fmt.Errorf("The package '%s' already exists", name)
	}

	if err := directory.Create(packagePath); err != nil {
		return nil, err
	}

	pckg, err := OpenPackage(packagePath)
	return pckg, err
}

// OpenPackage opens an existing package in a go repository
func (repository Repository) OpenPackage(name string) (*Package, error) {
	packagePath := path.Join(repository.path, name)
	pckg, err := OpenPackage(packagePath)

	if err != nil {
		return nil, err
	}

	return pckg, nil
}

// RemovePackage removes an existing package from the go repository
func (repository Repository) RemovePackage(name string) error {
	packagePath := path.Join(repository.path, name)
	if !directory.Exists(packagePath) {
		return fmt.Errorf("The package '%s' does not exist", name)
	}
	return directory.Remove(packagePath)
}

// RenamePackage renames an existing package in a go repository
func (repository Repository) RenamePackage(oldName string, newName string) (Package, error) {
	packagePath := path.Join(repository.path, oldName)
	pckg := Package{name: oldName, path: packagePath}
	newPackagePath := path.Join(repository.path, newName)

	if directory.Exists(newPackagePath) {
		return pckg, fmt.Errorf("The package '%s' already exists", newName)
	}
	err := directory.Rename(packagePath, newPackagePath)

	if err == nil {
		pckg.name = newName
		pckg.path = newPackagePath
	}

	return pckg, err
}
