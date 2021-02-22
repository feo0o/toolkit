package toolkit

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	t.Log(RandomString(0))
	t.Log(RandomString(8))
}

func TestRandomInt(t *testing.T) {
	for i := 0; i < 10; i++ {
		n, err := RandomInt(4, 19)
		if err != nil {
			t.Log(err)
			return
		}
		t.Log(n)
	}

	for i := 0; i < 10; i++ {
		n, err := RandomInt(6, 6)
		if err != nil {
			t.Log(err)
			return
		}
		t.Log(n)
	}

	for i := 0; i < 10; i++ {
		n, err := RandomInt(9, 4)
		if err != nil {
			t.Log(err)
			return
		}
		t.Log(n)
	}
}
