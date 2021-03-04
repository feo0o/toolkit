package toolkit

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestCreateFilePath(t *testing.T) {
	f := "test.log"
	if err := CreateFilePath(f); err != nil {
		t.Error(err)
	}
	if err := ioutil.WriteFile(f, []byte(fmt.Sprintf("%s: hello, there", time.Now().Format(time.RFC3339))), 0644); err != nil {
		t.Error(err)
	}
	time.Sleep(1 * time.Second)
	// test for recreate file
	if err := CreateFilePath(f); err != nil {
		t.Error(err)
	}
	b, err := ioutil.ReadFile(f)
	if err != nil {
		t.Error(err)
	}
	t.Log(time.Now().Format(time.RFC3339))
	t.Log(string(b))
}
