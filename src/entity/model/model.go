package model

import "io"
import "../../entity"
import "../../entity/builder"
import "../../bjson"

type Model interface {
	AttributeNames() [][]byte
	AttributeTypes() []byte
	Allocate([][]byte) Instance
}

type ModelDefinition struct {
	attributeNames [][]byte
	attributeTypes []byte
}

func (m *ModelDefinition) AttributeNames() [][]byte {
	return m.attributeNames
}
func (m *ModelDefinition) AttributeTypes() []byte {
	return m.attributeTypes
}

type Instance interface {
	Model() Model
	Id() int
	Write(io.Writer, byte)
	Update([][]byte)
}

func UpdateRecord(record *entity.Record, json []byte) {
	_, values := bjson.Parse(json, record.Definition.AttributeNames)
	builder.Update(record, values)
}
