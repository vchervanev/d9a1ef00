package builder

import "../../entity"

func Build(definition *entity.Definition, names, values [][]byte) (result *entity.Record) {
	record := entity.Record{Definition: definition, Data: values}
	return &record
}

func Update(record *entity.Record, names, values [][]byte) {
}
