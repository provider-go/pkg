package sm4

import "testing"

func TestName(t *testing.T) {
	instance := NewSMFour("gmsm")
	a := instance.Encode("1234")
	t.Log(a)
	b, _ := instance.Decode(a)
	t.Log(b)
}
