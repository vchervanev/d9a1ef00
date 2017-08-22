package storage

import "../../entity"

// MemoryService implements StorageService
type MemoryService struct {
	data Database
}

var MemoryServiceFactory MemoryServiceFactoryImpl

type MemoryServiceFactoryImpl struct{}

func (*MemoryServiceFactoryImpl) CreateMemoryService(typeNames []string) (result MemoryService) {
	result.data = make(Database, len(typeNames))
	for _, typeName := range typeNames {
		result.data[typeName] = make(RecordMap, 10000)
	}

	return
}

func (m *MemoryService) add(record entity.Record) {
	m.data[record.Definition.EntityType][record.Id()] = record
}

func (m *MemoryService) update(record entity.Record) {
	// TODO update new fields
}

func (m *MemoryService) get(entityType string, id int) entity.Record {
	return m.data[entityType][id]
}
