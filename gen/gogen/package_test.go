package gogen

import "testing"

func TestOpenPackage(t *testing.T) {
	pkg, _ := OpenPackage("../gogen")

	t.Log(pkg.name)
}
