package gogen

import "testing"

func TestOpenPackage(t *testing.T) {
	pkg, _ := NewPackage("../gogen")

	t.Log(pkg.name)
}
