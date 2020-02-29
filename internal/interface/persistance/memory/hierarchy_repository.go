package memory

import (
	"dast-api/internal/domain/model"
	"sync"
)

func NewHierarchyRepository() *hierarchyRepository {
	return &hierarchyRepository{
		mu:      &sync.Mutex{},
		durable: map[string]*model.Hierarchy{},
	}
}

type hierarchyRepository struct {
	mu      *sync.Mutex
	durable map[string]*model.Hierarchy
}

func (hr hierarchyRepository) Save(hierarchy *model.Hierarchy) error {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	hr.durable[hierarchy.ID] = hierarchy
	return nil
}

func (hr hierarchyRepository) Get(id string) (*model.Hierarchy, error) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	h, ok := hr.durable[id]
	if !ok {
		return nil, nil
	}
	return h, nil
}
