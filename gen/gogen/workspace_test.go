package gogen

import (
	"os"
	"path"
	"testing"
)

func getWorkspacePath() string {
	return path.Join(os.Getenv("GOPATH"), "src")
}

func TestCreateRepo(t *testing.T) {

	workspace := Workspace{name: "Test", path: getWorkspacePath()}

	_, error := workspace.CreateRepository("test")

	if error != nil {
		t.Error(error)
	}
}

func TestRemoveRepo(t *testing.T) {

	workspace := Workspace{name: "Test", path: getWorkspacePath()}

	error := workspace.RemoveRepository("test")

	if error != nil {
		t.Error(error)
	}
}
