package memory

import (
	"dast-api/internal/domain/model"
	"sync"
)

func NewTemplateRepository() *TemplateRepository {
	return &TemplateRepository{
		mu:      &sync.Mutex{},
		durable: map[string]*model.CriteriaTemplate{},
	}
}

type TemplateRepository struct {
	mu      *sync.Mutex
	durable map[string]*model.CriteriaTemplate
}

func (t TemplateRepository) Save(template *model.CriteriaTemplate) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.durable[template.ID] = template
	return nil
}

func (t TemplateRepository) SearchPublicTemplates() ([]*model.CriteriaTemplate, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	toRet := []*model.CriteriaTemplate{}
	for _, template := range t.durable {
		toRet = append(toRet, template)
	}
	return toRet, nil
}
