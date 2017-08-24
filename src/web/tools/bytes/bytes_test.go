package bytes

import (
	"testing"
)

func TestEql(t *testing.T) {
	b1 := []byte("aaa/bbb/")
	b2 := []byte("aaa/bbb/")
	if !Eql(b1, b2) {
		t.Error("expected true but got false")
	}

	b2 = []byte("aaa/bbb/c")
	if Eql(b1, b2) {
		t.Error("expected false but got true")
	}
}

func TestStartsWith(t *testing.T) { //bytes, prefix []byte) bool
	b1 := []byte("aaa/bbb/")
	prefix := []byte("aaa/b")

	if !StartsWith(b1, prefix) {
		t.Error("expected true but got false")
	}

	prefix = []byte("aaa/c")
	if StartsWith(b1, prefix) {
		t.Error("expected false but got true")
	}
}

func TestGetId(t *testing.T) { //bytes []byte, separator byte) int
	b1 := []byte("aaa/bbb/432")
	id := GetId(b1, '/')
	if id != 432 {
		t.Error("expected 432 but got", id)
	}

	b1 = []byte("aaa/bbb/")
	id = GetId(b1, '/')
	if id != 0 {
		t.Error("expected 0 but got", id)
	}
}
