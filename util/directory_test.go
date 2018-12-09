package directory

import (
	"testing"
)

func TestListFiles(t *testing.T) {

	files, _ := ListFiles("../", true)
	t.Log("test 0")
	t.Log(files)
	t.Error("test")
}
