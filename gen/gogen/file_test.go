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
	typeName2, typeDescription2 := "Test2", "This is a second test description"

	typ, err := file.AddType(typeName, typeDescription, "string")
	typ2, err := file.AddType(typeName2, typeDescription2, "int")

	test.CheckError(err, t)

	if typ.name != typeName {
		t.Fail()
	}
	if typ.description != typeDescription {
		t.Fail()
	}

	if typ2.name != typeName2 {
		t.Fail()
	}
	if typ2.description != typeDescription2 {
		t.Fail()
	}

	teardownFileTest(t)
}

func TestFile_RenameType(t *testing.T) {
	setupFileTest(t)

	workspace := Workspace{name: "Test", path: getFileTestPath()}
	repo, _ := workspace.CreateRepository("test-repo")
	pkg, _ := repo.CreatePackage("testpackage")
	file, _ := pkg.CreateFile("test1")

	typeName, typeDescription := "Test", "This is a test description"
	typeName2, typeDescription2 := "Test2", "This is a second test description"

	typ, err := file.AddType(typeName, typeDescription, "string")

	test.CheckError(err, t)

	err = file.RenameType(typeName, typeName2)

	test.CheckError(err, t)

	typ.SetDescription(typeDescription2)

	if typ.name != typeName2 {
		t.Fail()
	}
	if typ.description != typeDescription2 {
		t.Fail()
	}

	teardownFileTest(t)
}
