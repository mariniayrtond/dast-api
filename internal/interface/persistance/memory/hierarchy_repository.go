package memory

import (
	"dast-api/internal/domain/model"
	"sync"
)

func NewHierarchyRepository() *HierarchyRepository {
	return &HierarchyRepository{
		mu:      &sync.Mutex{},
		durable: map[string]*model.Hierarchy{},
	}
}

type HierarchyRepository struct {
	mu      *sync.Mutex
	durable map[string]*model.Hierarchy
}

func (hr HierarchyRepository) Save(hierarchy *model.Hierarchy) error {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	hr.durable[hierarchy.ID] = hierarchy
	return nil
}

func (hr HierarchyRepository) Get(id string) (*model.Hierarchy, error) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	h, ok := hr.durable[id]
	if !ok {
		return nil, nil
	}
	return h, nil
}

func (hr HierarchyRepository) SearchByUsername(username string) ([]*model.Hierarchy, error) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	toRet := []*model.Hierarchy{}
	for _, h := range hr.durable {
		if h.Owner == username {
			toRet = append(toRet, h)
		}
	}

	return toRet, nil
}
