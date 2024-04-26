package util

import "testing"

func Test_Unzip(t *testing.T) {
	err := Unzip("demo.zip", "./")
	if err != nil {
		t.Log(err)
	}
}
