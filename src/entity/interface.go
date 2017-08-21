package entity

type Definition struct {
	EntityName     string
	AttributeNames [][]byte
}

type Record struct {
	Definition *Definition
	Data       [][]byte
}
