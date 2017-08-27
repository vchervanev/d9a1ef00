package visit

import "../../../entity"
import "../../builder"
import "../../../bjson"

var ID = []byte("id")
var LOCATION = []byte("location")
var USER = []byte("user")
var VISITED_AT = []byte("visited_at")
var MARK = []byte("mark")

var VisitDefinition = entity.Definition{
	EntityType:     "visit",
	AttributeNames: [][]byte{ID, LOCATION, USER, VISITED_AT, MARK},
	AttributeTypes: []byte{'n', 'n', 'n', 'n', 'n'},
}

func BuildVisit(json []byte) *entity.Record {
	var _, values = bjson.Parse(json, VisitDefinition.AttributeNames)
	record := builder.Build(&VisitDefinition, values)
	return record
}
