package memory

import (
	"dast-api/internal/domain/model"
	"sync"
)

func NewCriteriaJudgementsRepository() *CriteriaJudgementsRepository {
	return &CriteriaJudgementsRepository{
		mu:      &sync.Mutex{},
		durable: map[string]*model.CriteriaJudgements{},
	}
}

type CriteriaJudgementsRepository struct {
	mu      *sync.Mutex
	durable map[string]*model.CriteriaJudgements
}

func (hr CriteriaJudgementsRepository) Save(pWise *model.CriteriaJudgements) error {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	hr.durable[pWise.ID] = pWise
	return nil
}

func (hr CriteriaJudgementsRepository) Get(id string) (*model.CriteriaJudgements, error) {
	hr.mu.Lock()
	defer hr.mu.Unlock()

	h, ok := hr.durable[id]
	if !ok {
		return nil, nil
	}
	return h, nil
}
