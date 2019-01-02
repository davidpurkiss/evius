package gogen

import (
	"evius/util/directory"
	"evius/util/test"
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

	workspace := Workspace{name: "Test", path: getRepositoryTestPath()}
	repo, err := workspace.CreateRepository("test-repo")

	test.CheckError(err, t)

	pkg, err := repo.CreatePackage("test-package")

	test.CheckError(err, t)

	if !directory.Exists(pkg.path) {
		t.Fail()
	}

	teardownRepositoryTest(t)
}

func TestRemovePackage(t *testing.T) {

	setupRepositoryTest(t)

	workspace := Workspace{name: "Test", path: getRepositoryTestPath()}
	repo, err := workspace.CreateRepository("test-repo")
	test.CheckError(err, t)

	pkg, err := repo.CreatePackage("test-package")

	test.CheckError(err, t)

	test.CheckError(repo.RemovePackage(pkg.name), t)

	if directory.Exists(pkg.path) {
		t.Fail()
	}

	teardownRepositoryTest(t)
}

func TestRenamePackage(t *testing.T) {

	setupRepositoryTest(t)

	workspace := Workspace{name: "Test", path: getRepositoryTestPath()}
	repo, err := workspace.CreateRepository("test-repo")
	test.CheckError(err, t)

	pkg, err := repo.CreatePackage("test-package")
	test.CheckError(err, t)

	repo.RenamePackage(pkg.name, "test-package-renamed")

	if directory.Exists(pkg.name) {
		t.Fail()
	}

	if directory.Exists("test-package") {
		t.Fail()
	}

	teardownRepositoryTest(t)
}

func TestOpenPackage(t *testing.T) {

	workspace := Workspace{name: "Test", path: test.GetTestDataPath()}
	repo, err := workspace.OpenRepository(getRepositoryTestName())

	test.CheckError(err, t)

	pkg, err := repo.OpenPackage("testpackage")

	test.CheckError(err, t)

}
