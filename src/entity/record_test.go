package entity

import "testing"

var definition = Definition{
	EntityType:     "test",
	AttributeNames: [][]byte{[]byte("id"), []byte("attribute")},
}

var record = Record{
	Definition: &definition,
	Data:       [][]byte{[]byte("5432"), []byte("value")},
}

func TestId(t *testing.T) {
	const expected = 5432
	actual := record.Id()
	if actual != expected {
		t.Error("Expected ", expected, "but got", actual)
	}
}
