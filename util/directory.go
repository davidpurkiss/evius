package directory

import "os"

// Create creates a new directory
func Create(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Remove removes an existing directory
func Remove(path string) error {
	return os.Remove(path)
}

// Rename renames an existing directory
func Rename(oldPath string, newPath string) error {
	return os.Rename(oldPath, newPath)
}

// Exists checks to see if a directory exists or not
func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
