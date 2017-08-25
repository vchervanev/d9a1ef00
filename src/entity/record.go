package entity

import "io"

type Definition struct {
	EntityType     string
	AttributeNames [][]byte
	AttributeTypes []byte
}

type Record struct {
	Definition *Definition
	Data       [][]byte
}

func (record *Record) Id() (result int) {
	result = 0
	for _, d := range record.Data[0] {
		result = result*10 + int(d-'0')
	}
	return
}

func (record *Record) WriteJSON(dst io.Writer) {
	dst.Write([]byte{'{'})

	for i, name := range record.Definition.AttributeNames {
		if i != 0 {
			dst.Write([]byte{','})
		}
		dst.Write([]byte{'"'})
		dst.Write(name)
		if record.Definition.AttributeTypes[i] == 's' {
			dst.Write([]byte("\":\""))
		} else {
			dst.Write([]byte("\":"))
		}
		dst.Write(record.Data[i])
		if record.Definition.AttributeTypes[i] == 's' {
			dst.Write([]byte{'"'})
		}
	}
	dst.Write([]byte{'}'})

}
