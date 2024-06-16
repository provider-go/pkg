package did

import (
	"testing"
	"time"
)

func TestAAA(t *testing.T) {
	//to, _ := time.Parse("2006-01-02 15:04:05", "2021-01-27 10:10:10.294")
	stamp := time.Now().Format("2006-01-02T15:04:05Z")
	t.Log(stamp)
}

func TestCreateDIDDocument(t *testing.T) {
	document := CreateDIDDocument("3737d7558851a6158fd95d7748ac49ffbc20b9d8507534de82cefce11223b9c3", "123456")
	t.Log(document)
}

func TestCreateVCDocument(t *testing.T) {
	info := "{ \"name\" :\"王奇\"}"
	document := CreateVCDocument("3737d7558851a6158fd95d7748ac49ffbc20b9d8507534de82cefce11223b9c3", "123456", "13", "did:cmid:AAA", "365", "did:cmid:BBB", info)
	t.Log(document)
}

func TestCreateVPDocument(t *testing.T) {
	info := "{ \"name\" :\"王奇\"}"
	vc := "[" + CreateVCDocument("3737d7558851a6158fd95d7748ac49ffbc20b9d8507534de82cefce11223b9c3", "123456", "13", "did:cmid:AAA", "365", "did:cmid:BBB", info) + "]"
	document := CreateVPDocument("3737d7558851a6158fd95d7748ac49ffbc20b9d8507534de82cefce11223b9c3", "123456", "did:cmid:AAA", vc)
	t.Log(document)
}
