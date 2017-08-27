package builder

import "../../entity"
import "../../tools/bytes"

func Build(definition *entity.Definition, names, values [][]byte) (result *entity.Record) {
	data := make([][]byte, len(values))
	count := len(names)
	for i := 0; i < count; i++ {
		// check values order:
		j := bytes.IndexOf(definition.AttributeNames, names[i])
		data[j] = values[i]
	}
	record := entity.Record{Definition: definition, Data: data}
	return &record
}

func Update(record *entity.Record, names, values [][]byte) {
}
