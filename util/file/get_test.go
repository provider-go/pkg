package file

import "testing"

func Test_ListDir(t *testing.T) {
	dir, err := ListDir("../../../plugin")
	if err != nil {
		t.Log(err)
	}
	t.Log(dir)
}
