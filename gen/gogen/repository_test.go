package gogen

import (
	"evius/util/directory"
	"evius/util/test"
	"path"
	"testing"
)

func getRepositoryTestName() string {
	return "repository"
}

func setupRepositoryTest(t *testing.T) {
	test.Setup(getRepositoryTestName(), t)
}

func teardownRepositoryTest(t *testing.T) {
	test.Teardown(getRepositoryTestName(), t)
}

func getRepositoryTestPath() string {
	return test.GetTestPath(getRepositoryTestName())
}

func TestCreatePackage(t *testing.T) {

	setupRepositoryTest(t)

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
