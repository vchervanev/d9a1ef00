package zip

import "testing"

func expect(text string, value int, actual int, t *testing.T) {
	if value != actual {
		t.Error("expected", text, "to be", value, "but got", actual)
	}

}
func TestLoader(t *testing.T) {
	stat := make(map[byte]int)
	handler := func(record []byte, filter byte) {
		stat[filter] += 1
	}

	LoadObjects("../../data/data.zip", []byte("ulv"), handler)

	expect("object count", 3, len(stat), t)
	expect("u objects", 1091, stat['u'], t)
	expect("l objects", 828, stat['l'], t)
	expect("v objects", 10910, stat['v'], t)
}
