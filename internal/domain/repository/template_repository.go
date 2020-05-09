package repository

import "dast-api/internal/domain/model"

type TemplateRepository interface {
	Save(hierarchy *model.CriteriaTemplate) error
	SearchPublicTemplates() ([]*model.CriteriaTemplate, error)
}
