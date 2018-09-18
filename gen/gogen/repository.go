package gogen

import (
	"evius/util"
	"fmt"
	"path"
)

// Repository defines the attributes of a go repository
type Repository struct {
	name string
	path string
}

// CreatePackage creates a new package in the go repository
func (repository Repository) CreatePackage(name string) (Package, error) {
	packagePath := path.Join(repository.path, name)
	if directory.Exists(packagePath) {
		return Package{}, fmt.Errorf("The package '%s' already exists", name)
	}
	pckg := Package{name: name, path: packagePath}
	return pckg, directory.Create(pckg.path)
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
