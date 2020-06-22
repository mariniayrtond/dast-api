package repository

import "dast-api/internal/domain/model"

//go:generate moq -out hierarchy_repository_moq.go . HierarchyRepository

type HierarchyRepository interface {
	Override(id string, hierarchy *model.Hierarchy) error
	Save(hierarchy *model.Hierarchy) error
	Get(id string) (*model.Hierarchy, error)
	SearchByUsername(value string) ([]*model.Hierarchy, error)
}
