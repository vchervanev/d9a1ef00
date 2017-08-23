package user

import "../../../entity"
import "../../builder"
import "../../../bjson"

var ID = []byte("id")

var UserDefinition = entity.Definition{
	EntityType:     "user",
	AttributeNames: [][]byte{ID},
}

// func Id(record *entity.Record) []byte {
// 	return record.Data[0]
// }

func BuildUser(json []byte) *entity.Record {
	var names, values = bjson.Parse(json)
	record := builder.Build(&UserDefinition, names, values)
	return record
}
