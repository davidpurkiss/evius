package test

import (
	"evius/util/directory"
	"os"
	"path"
	"testing"
)

// GetGoWorkspacePath returns the path of the src directory in the GOPATH
func GetGoWorkspacePath() string {
	return path.Join(os.Getenv("GOPATH"), "src")
}

// GetTestPath return the path for the specific set of tests, e.g. workspace, repo, package, file, etc.
func GetTestPath(testName string) string {
	return path.Join(GetTestRootPath(), testName)
}

// GetTestRootPath returns the root path for tests in the gogen package
func GetTestRootPath() string {
	return path.Join(GetGoWorkspacePath(), "evius-test", "gen", "gogen")
}

// Setup performs base setup for tests, creating test directory
func Setup(testName string, t *testing.T) {

	if directory.Exists(GetTestPath(testName)) {
		if err := directory.RemoveAll(GetTestPath(testName)); err != nil {
			t.Error(err)
		}
	}

	if err := directory.Create(GetTestPath(testName)); err != nil {
		t.Error(err)
	}
}

// Teardown performs base teardown for test, removing test directory and all contents
func Teardown(testName string, t *testing.T) {
	if err := directory.RemoveAll(GetTestPath(testName)); err != nil {
		t.Error(err)
	}
}
