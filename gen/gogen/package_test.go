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

	_, err = pkg.CreateFile("test1")

	test.CheckError(err, t)

	err = pkg.RemoveFile("test1")

	test.CheckError(err, t)

	if len(pkg.files) != 0 {
		t.Fail()
	}

	teardownPackageTest(t)
}

func TestRenameFile(t *testing.T) {

	setupPackageTest(t)

	workspace := Workspace{name: "Test", path: getPackageTestPath()}
	repo, err := workspace.CreateRepository("test-repo")
	test.CheckError(err, t)

	pkg, err := repo.CreatePackage("testpackage")

	test.CheckError(err, t)

	file, err := pkg.CreateFile("test1")

	test.CheckError(err, t)

	file1Path := pkg.GetFilePath("test1")
	file2Path := pkg.GetFilePath("test2")

	_, err = pkg.RenameFile("test1", "test2")

	if directory.Exists(file1Path) {
		t.Fail()
	}

	if !directory.Exists(file2Path) {
		t.Fail()
	}

	if testFile := pkg.GetFile("test1"); testFile != nil {
		t.Fail()
	}

	if testFile := pkg.GetFile("test2"); testFile == nil {
		t.Fail()
	}

	if file.name != "test2" && file.path != file2Path {
		t.Fail()
	}

	teardownPackageTest(t)
}

func TestOpenFile(t *testing.T) {

	workspace := Workspace{name: "Test", path: path.Join(test.GetTestDataPath(), "repository")}
	repo, err := workspace.OpenRepository("testrepo")

	test.CheckError(err, t)

	pkg, err := repo.OpenPackage("testpackage")

	test.CheckError(err, t)

	if pkg != nil {
		if len(pkg.files) != 2 {
			t.Error("Expected 2 files")
		}
	}

}
