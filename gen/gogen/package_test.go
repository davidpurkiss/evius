package gogen

import (
	"evius/util/directory"
	"evius/util/test"
	"path"
	"testing"
)

func getPackageTestName() string {
	return "package"
}

func setupPackageTest(t *testing.T) {
	test.Setup(getPackageTestName(), t)
}

func teardownPackageTest(t *testing.T) {
	test.Teardown(getPackageTestName(), t)
}

func getPackageTestPath() string {
	return test.GetTestPath(getPackageTestName())
}

func TestCreateFile(t *testing.T) {

	setupPackageTest(t)

	workspace := Workspace{name: "Test", path: getPackageTestPath()}
	repo, err := workspace.CreateRepository("test-repo")

	test.CheckError(err, t)

	pkg, err := repo.CreatePackage("testpackage")

	test.CheckError(err, t)

	file, err := pkg.CreateFile("test1")

	test.CheckError(err, t)

	if file == nil {
		t.Error("File not defined")
	}

	if file.name != "test1" {
		t.Error("File does not have the correct name")
	}

	if file.path != path.Join(pkg.path, "test1.go") {
		t.Error("File does not have the correct path")
	}

	if file._file == nil {
		t.Error("Ast file not defined")
	}

	teardownPackageTest(t)
}

func TestRemoveFile(t *testing.T) {

	setupPackageTest(t)

	workspace := Workspace{name: "Test", path: getPackageTestPath()}
	repo, err := workspace.CreateRepository("test-repo")
	test.CheckError(err, t)

	pkg, err := repo.CreatePackage("testpackage")

	test.CheckError(err, t)

	test.CheckError(repo.RemovePackage(pkg.name), t)

	if directory.Exists(pkg.path) {
		t.Fail()
	}

	teardownPackageTest(t)
}

func TestRenameFile(t *testing.T) {

	setupPackageTest(t)

	workspace := Workspace{name: "Test", path: getPackageTestPath()}
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

	teardownPackageTest(t)
}

func TestOpenFile(t *testing.T) {

	workspace := Workspace{name: "Test", path: path.Join(test.GetTestDataPath(), "repository")}
	repo, err := workspace.OpenRepository(getPackageTestName())

	test.CheckError(err, t)

	pkg, err := repo.OpenPackage("testpackage")

	test.CheckError(err, t)

	if pkg != nil {
		if len(pkg.files) != 2 {
			t.Error("Expected 2 files")
		}
	}

}
