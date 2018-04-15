package storage

import (
	"fmt"
	"sync"

	"../model"
)

type RecordMap map[int]*model.Instance
type Database map[model.Model]RecordMap

// MemoryService implements StorageService
type MemoryService struct {
	data Database
}

var MemoryServiceFactory MemoryServiceFactoryImpl

type MemoryServiceFactoryImpl struct{}

func (*MemoryServiceFactoryImpl) CreateMemoryService(models []model.Model, sizes []int) (result MemoryService) {
	result.data = make(Database, len(models))
	for i, model := range models {
		result.data[model] = make(RecordMap, sizes[i])
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

func (m *MemoryService) Add(record *model.Instance) {
	writeMutex.Lock()
	model := (*record).Model()
	id := (*record).Id()
	m.data[model][id] = record
	writeMutex.Unlock()
}

// func (m *MemoryService) Update(record entity.Record) {
// 	// TODO update new fields
// }

func (m *MemoryService) Get(model *model.Model, id int) *model.Instance {
	return m.data[*model][id]
}
