package storage

import "../../entity"

type RecordMap map[int]*entity.Record
type Database map[string]RecordMap

type StorageService interface {
	add(record entity.Record)
	update(record entity.Record)
	get(entityType string, id int) entity.Record
}
