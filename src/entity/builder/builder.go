package builder

import "../../entity"

func Build(definition *entity.Definition, names, values [][]byte) (result *entity.Record) {
	record := entity.Record{Definition: definition, Data: values}
	// TODO check values order!
	return &record
}

func Update(record *entity.Record, names, values [][]byte) {
}
