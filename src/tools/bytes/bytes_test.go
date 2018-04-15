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

func TestStartsWith(t *testing.T) {
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

func TestGetId(t *testing.T) {
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

func TestIndexOfPositive(t *testing.T) {
	item1 := []byte("value1")
	item2 := []byte("value2")
	item3 := []byte("value3")
	target := []byte("value2")
	items := [][]byte{item1, item2, item3}

	result := IndexOf(items, target)
	if result != 1 {
		t.Error("Expected 1 but got", result)
	}
}

func TestIndexOfNegative(t *testing.T) {
	item1 := []byte("value1")
	item2 := []byte("value2")
	item3 := []byte("value3")
	target := []byte("value9")
	items := [][]byte{item1, item2, item3}

	result := IndexOf(items, target)
	if result != -1 {
		t.Error("Expected -1 but got", result)
	}
}
func TestIndexOfSubarray(t *testing.T) {
	source := []byte("initialvalueishere")
	value := []byte("value")

	result := IndexOfSubarray(source, value)
	if result != 7 {
		t.Error("Expected 7 but got", result)
	}
}

func TestIndexOfSubarrayNegative(t *testing.T) {
	source := []byte("initialvalueishere")
	value := []byte("value1")

	result := IndexOfSubarray(source, value)
	if result != -1 {
		t.Error("Expected -1 but got", result)
	}
}
