package repository

import "dast-api/internal/domain/model"

//go:generate moq -out hierarchy_repository_moq.go . HierarchyRepository

type HierarchyRepository interface {
	Save(hierarchy *model.Hierarchy) error
	Get(id string) (*model.Hierarchy, error)
}
