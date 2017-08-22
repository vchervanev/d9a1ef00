package entity

type Definition struct {
	EntityType     string
	AttributeNames [][]byte
}

type Record struct {
	Definition *Definition
	Data       [][]byte
}
