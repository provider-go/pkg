package gmsm

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	fmt.Println("ret = ", ret)
	return ret
}
func TestSm3(t *testing.T) {
	msg := []byte("test")
	err := ioutil.WriteFile("ifile", msg, os.FileMode(0644)) // 生成测试文件
	if err != nil {
		t.Fatal(err)
	}
	msg, err = ioutil.ReadFile("ifile")
	if err != nil {
		t.Fatal(err)
	}
	hw := New()
	hw.Write(msg)
	hash := hw.Sum(nil)
	fmt.Println(hash)
	fmt.Printf("hash = %d\n", len(hash))
	fmt.Printf("%s\n", byteToString(hash))
	hash1 := Sm3Sum(msg)
	fmt.Println(hash1)
	fmt.Printf("%s\n", byteToString(hash1))

}
