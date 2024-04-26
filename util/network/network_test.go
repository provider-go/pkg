package util

import "testing"

func Test_GetUseIp(t *testing.T) {
	addr := GetHostIp()
	t.Log(addr)
}
