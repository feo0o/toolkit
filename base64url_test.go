package toolkit

import "testing"

var src = "+/="

func TestBase64URLEncode(t *testing.T) {
	t.Log(Base64URLEncode([]byte(src)))
}

func TestBase64URLDecode(t *testing.T) {
	s, err := Base64URLDecode("Ky89")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(s))
}
