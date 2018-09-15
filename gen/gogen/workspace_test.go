package gogen

import (
	"os"
	"path"
	"testing"
)

func TestCreateWorkspace(t *testing.T) {

	workspacePath := path.Join(os.Getenv("GOPATH"), "src", "test")
	workspace := Workspace{name: "Test", path: workspacePath}

	workspace.CreateWorkspace()
}
