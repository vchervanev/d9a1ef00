package entity

type Definition struct {
	EntityType     string
	AttributeNames [][]byte
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
