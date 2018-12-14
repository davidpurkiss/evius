package gogen

import "testing"

func TestOpenPackage(t *testing.T) {
	pkg, _ := OpenPackage("../gogen")
	file, _ := pkg.CreateFile("test")
	t.Log(pkg.name)
	t.Log(file.name)
}
