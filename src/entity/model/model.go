package model

import "../../entity"
import "../../entity/builder"
import "../../bjson"

func UpdateRecord(record *entity.Record, json []byte) {
	_, values := bjson.Parse(json, record.Definition.AttributeNames)
	builder.Update(record, values)
}
