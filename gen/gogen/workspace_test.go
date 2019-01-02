package gogen

import (
	"evius/util/directory"
	"evius/util/test"
	"path"
	"testing"
)

func getWorkspaceTestName() string {
	return "workspace"
}

func setupWorkspaceTest(t *testing.T) {
	test.Setup(getWorkspaceTestName(), t)
}

func teardownWorkspaceTest(t *testing.T) {
	test.Teardown(getWorkspaceTestName(), t)
}

func getWorkspaceTestPath() string {
	return test.GetTestPath(getWorkspaceTestName())
}

func TestCreateRepo(t *testing.T) {

	setupWorkspaceTest(t)

	workspace := Workspace{name: "Test", path: getWorkspaceTestPath()}

	repo, err := workspace.CreateRepository("test")

	if err != nil {
		t.Error(err)
	}

	if !directory.Exists(repo.path) {
		t.Fail()
	}

	teardownWorkspaceTest(t)
}

func TestRemoveRepo(t *testing.T) {

	setupWorkspaceTest(t)

	workspace := Workspace{name: "Test", path: getWorkspaceTestPath()}

	repo, err := workspace.CreateRepository("test")

	if err != nil {
		t.Error(err)
	}

	err = workspace.RemoveRepository("test")

	if err != nil {
		t.Error(err)
	}

	if directory.Exists(repo.path) {
		t.Fail()
	}

	teardownWorkspaceTest(t)
}

func TestRenameRepo(t *testing.T) {

	setupWorkspaceTest(t)

	workspace := Workspace{name: "Test", path: getWorkspaceTestPath()}

	_, err := workspace.CreateRepository("test")

	repo, err := workspace.RenameRepository("test", "test2")

	if err != nil {
		t.Error(err)
	}

	if !directory.Exists(repo.path) {
		t.Fail()
	}

	if directory.Exists(path.Join(getWorkspaceTestPath(), "test")) {
		t.Fail()
	}

	teardownWorkspaceTest(t)
}
