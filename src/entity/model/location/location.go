package location

import "../../../entity"
import "../../builder"
import "../../../bjson"

var ID = []byte("id")
var PLACE = []byte("place")
var COUNTRY = []byte("country")
var CITY = []byte("city")
var DISTANCE = []byte("distance")

var LocationDefinition = entity.Definition{
	EntityType:     "location",
	AttributeNames: [][]byte{ID, PLACE, COUNTRY, CITY, DISTANCE},
	AttributeTypes: []byte{'n', 's', 's', 's', 'n'},
}

func BuildLocation(json []byte) *entity.Record {
	var names, values = bjson.Parse(json, LocationDefinition.AttributeNames)
	record := builder.Build(&LocationDefinition, names, values)
	return record
}
