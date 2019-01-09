package gogen

import (
	"evius/util/test"
	"testing"
)

func getFileTestName() string {
	return "file"
}

func setupFileTest(t *testing.T) {
	test.Setup(getFileTestName(), t)
}

func teardownFileTest(t *testing.T) {
	test.Teardown(getFileTestName(), t)
}

func getFileTestPath() string {
	return test.GetTestPath(getFileTestName())
}

func TestFile_AddType(t *testing.T) {
	setupFileTest(t)

	workspace := Workspace{name: "Test", path: getFileTestPath()}
	repo, _ := workspace.CreateRepository("test-repo")
	pkg, _ := repo.CreatePackage("testpackage")
	file, _ := pkg.CreateFile("test1")

	typeName, typeDescription := "Test", "This is a test description"

	typ, err := file.AddType(typeName, typeDescription, "string")

	test.CheckError(err, t)

	if typ.name != typeName {
		t.Fail()
	}
	if typ.description != typeDescription {
		t.Fail()
	}

	teardownFileTest(t)
}
