package repository

import "dast-api/internal/domain/model"

type HierarchyRepository interface {
	Save(hierarchy *model.Hierarchy) error
	Get(id string) (*model.Hierarchy, error)
}
