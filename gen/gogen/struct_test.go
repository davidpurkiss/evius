package gogen

import (
	"evius/util/test"
	"testing"
)

func getStructTestName() string {
	return "struct"
}

func setupStructTest(t *testing.T) {
	test.Setup(getStructTestName(), t)
}

func teardownStructTest(t *testing.T) {
	test.Teardown(getStructTestName(), t)
}

func getStructTestPath() string {
	return test.GetTestPath(getStructTestName())
}

func TestStruct_AddField(t *testing.T) {
	setupStructTest(t)

	workspace := Workspace{name: "Test", path: getStructTestPath()}
	repo, _ := workspace.CreateRepository("test-repo")
	pkg, _ := repo.CreatePackage("testpackage")
	file, _ := pkg.CreateFile("test1")

	structName, structDescription := "Test", "This is a test description"
	fieldName, fieldDescription, fieldType := "TestField1", "This is a string test field 1", "string"
	fieldName2, fieldDescription2, fieldType2 := "TestField2", "This is a string test field 2", "int"

	strct, err := file.AddStruct(structName, structDescription)

	test.CheckError(err, t)

	field, err := strct.AddField(fieldName, fieldType, fieldDescription)

	_, err = strct.AddField(fieldName2, fieldType2, fieldDescription2)

	if field.name != structName {
		t.Fail()
	}
	if field.description != structDescription {
		t.Fail()
	}
	if field.typeName != fieldType {
		t.Fail()
	}
	teardownStructTest(t)
}
