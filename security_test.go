package toolkit

import "testing"

var k = "LcYb9tNY97Kin28C"

func TestAESEncryptWithCFB(t *testing.T) {
	s := "hello world"
	c, err := AESEncryptWithCFB(s, k)
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}

func TestAESDecryptWithCFB(t *testing.T) {
	c := "494d0f74f5f887f6259eaa0f963c179211eabad7142260ff260f2e"
	s, err := AESDecryptWithCFB(c, k)
	if err != nil {
		t.Log(err)
	}
	t.Log(s)
}
