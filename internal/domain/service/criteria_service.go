package service

import (
	"dast-api/internal/domain/repository"
)

type CriteriaService struct {
	repo repository.HierarchyRepository
}

func NewCriteriaService(r repository.HierarchyRepository) *CriteriaService {
	return &CriteriaService{r}
}