package directory

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Create creates a new directory
func Create(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Remove removes an existing directory
func Remove(path string) error {
	return os.Remove(path)
}

// RemoveAll removes an existing directory and all its contents
func RemoveAll(path string) error {
	return os.RemoveAll(path)
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

// ListFiles returns a list of files in the specified directory.
func ListFiles(path string, recurse bool) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(path)

	if err != nil {
		return nil, err
	}

	files := make([]string, 0)

	for _, f := range fileInfos {

		if f.Name() == ".git" {
			continue
		}

		absoluteFilePath := getAbsoluteFilePathFromFileNameAndPath(path, f.Name())

		if !f.IsDir() {
			files = append(files, absoluteFilePath)
		} else if recurse {

			nestedFiles := make([]string, 0)
			nestedFiles, err = ListFiles(absoluteFilePath, recurse)

			if err != nil {
				return nil, err
			}

			for _, nf := range nestedFiles {
				files = append(files, nf)
			}
		}
	}

	return files, nil
}

func getAbsoluteFilePathFromFileNameAndPath(path string, fileName string) string {
	absolutePath, _ := filepath.Abs(path)
	return filepath.Join(absolutePath, fileName)
}
