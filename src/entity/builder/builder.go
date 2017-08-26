package builder

import "../../entity"

func dup(src []byte) (result []byte) {
	result = make([]byte, len(src))
	copy(result, src)
	return
}

func Build(definition *entity.Definition, names, values [][]byte) (result *entity.Record) {
	data := make([][]byte, len(values))
	count := len(names)
	for i := 0; i < count; i++ {
		data[i] = dup(values[i])
	}
	record := entity.Record{Definition: definition, Data: data}
	return &record
}

func Update(record *entity.Record, names, values [][]byte) {
}
