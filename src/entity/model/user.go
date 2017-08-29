package model

// import "../../../entity"
// import "../../builder"
// import "../../../bjson"

// var ID = []byte("id")
// var EMAIL = []byte("email")
// var FIRST_NAME = []byte("first_name")
// var LAST_NAME = []byte("last_name")
// var GENDER = []byte("gender")
// var BIRTH_DATE = []byte("birth_date")

// var UserDefinition = entity.Definition{
// 	EntityType:     "user",
// 	AttributeNames: [][]byte{ID, EMAIL, FIRST_NAME, LAST_NAME, GENDER, BIRTH_DATE},
// 	AttributeTypes: []byte{'n', 's', 's', 's', 's', 'n'},
// }

// func BuildUser(json []byte) *entity.Record {
// 	var _, values = bjson.Parse(json, UserDefinition.AttributeNames)
// 	record := builder.Build(&UserDefinition, values)
// 	return record
// }
