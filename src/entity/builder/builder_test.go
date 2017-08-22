package builder

import "testing"
import "../../entity"

// func Build(definition *entity.Definition, names, values [][]byte) (result entity.Record) {}

// func Update(record *entity.Record, names, values [][]byte) {}

func bytesEql(bytes1, bytes2 []byte, t *testing.T) {
	s1 := string(bytes1)
	s2 := string(bytes2)
	if s1 != s2 {
		t.Error("Byte sequences are not equal:", s1, "vs", s2)
	}
}

var attr1 = []byte("attr1")
var attr2 = []byte("attr2")
var attr3 = []byte("attr3")

var value1 = []byte("value1")
var value2 = []byte("value2")
var value3 = []byte("value3")
var value4 = []byte("value4")
var value5 = []byte("value5")

var definition = &entity.Definition{
	EntityType:     "test",
	AttributeNames: [][]byte{attr1, attr2, attr3},
}

func TestBuild(t *testing.T) {
	record := Build(definition, [][]byte{attr3, attr2, attr1}, [][]byte{value3, value2, value1})

	if len(record.Data) != 3 {
		t.Fatal("Expected array of 3 records")
	}
	bytesEql(record.Data[0], value1, t)
	bytesEql(record.Data[1], value2, t)
	bytesEql(record.Data[2], value3, t)
}

func TestUpdate(t *testing.T) {
	record := entity.Record{Definition: definition, Data: [][]byte{value1, value2, value3}}

	Update(&record, [][]byte{attr3, attr1}, [][]byte{value4, value5})

	if len(record.Data) != 3 {
		t.Fatal("Expected array of 3 records")
	}
	bytesEql(record.Data[0], value5, t)
	bytesEql(record.Data[1], value2, t)
	bytesEql(record.Data[2], value4, t)
}
