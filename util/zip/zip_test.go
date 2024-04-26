package util

import "testing"

func TestZip(t *testing.T) {
	err := Zip("qiqi.zip", "../math", "../net")
	if err != nil {
		t.Log(err)
	}
}
