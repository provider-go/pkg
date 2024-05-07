package util

import "testing"

func TestGetRandString(t *testing.T) {
	t.Log(GetRandString(16))
}

func TestGetMd5String(t *testing.T) {
	t.Log(GetMd5String("abcdef"))
}
