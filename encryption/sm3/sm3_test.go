package sm3

import "testing"

func TestName(t *testing.T) {
	instance := NewSMThree("gmsm")
	a := instance.Hash([]byte("aaa"))
	t.Log(a)

}
