package gogen

import "testing"

func TestCreatePackage(t *testing.T) {

	pkg, _ := OpenPackage("../../../test")
	file, _ := pkg.CreateFile("testfile")
	file2, _ := pkg.CreateFile("testfile2")
	t.Log(pkg.name)
	t.Log(file.name)
	t.Log(file2.name)
}
