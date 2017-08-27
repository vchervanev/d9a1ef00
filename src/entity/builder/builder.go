package builder

import "../../entity"

func Build(definition *entity.Definition, values [][]byte) (result *entity.Record) {
	record := entity.Record{Definition: definition, Data: values}
	return &record
}

func Update(record *entity.Record, values [][]byte) {
	for i, value := range values {
		if value != nil {
			record.Data[i] = value
		}
	}
}
