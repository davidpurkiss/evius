package gogen

import "testing"

func TestOpenPackage(t *testing.T) {
	pkg, _ := OpenPackage("../gogen")
	file, _ := pkg.CreateFile("test")
	testType, _ := file.AddType("Test", "This is a test type", "string")
	file.Save()
	t.Log(pkg.name)
	t.Log(file.name)
	t.Log(testType.name)
}
