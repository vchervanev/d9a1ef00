package storage

import (
	"fmt"
	"sync"

	"../../entity"
)

type RecordMap map[int]*entity.Record
type Database map[string]RecordMap

// MemoryService implements StorageService
type MemoryService struct {
	data Database
}

var MemoryServiceFactory MemoryServiceFactoryImpl

type MemoryServiceFactoryImpl struct{}

func (*MemoryServiceFactoryImpl) CreateMemoryService(typeNames []string, sizes []int) (result MemoryService) {
	result.data = make(Database, len(typeNames))
	for i, typeName := range typeNames {
		result.data[typeName] = make(RecordMap, sizes[i])
	}

	return
}

func (m *MemoryService) Info() string {
	stat := ""
	for key, records := range m.data {
		stat += fmt.Sprintf("%s: %v ", key, len(records))
	}
	return fmt.Sprintf("DB INFO[ %v]", stat)
}

var writeMutex sync.Mutex

func (m *MemoryService) Add(record *entity.Record) {
	writeMutex.Lock()
	m.data[record.Definition.EntityType][record.Id()] = record
	writeMutex.Unlock()
}

func (m *MemoryService) Update(record entity.Record) {
	// TODO update new fields
}

func (m *MemoryService) Get(entityType string, id int) *entity.Record {
	return m.data[entityType][id]
}
