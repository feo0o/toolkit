package toolkit

import "testing"

func TestNowCST(t *testing.T) {
	t.Log(NowCST())
}

func TestNowDateCST(t *testing.T) {
	t.Log(NowDateCST())
}

func TestNowCSTUnix(t *testing.T) {
	t.Log(NowCSTUnix())
}
