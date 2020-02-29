package usecase

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/domain/repository"
	"dast-api/internal/domain/service"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type HierarchyCRUD interface {
	RegisterHierarchy(string, string, string, []string, string) (*model.Hierarchy, error)
	GetHierarchy(id string) (*model.Hierarchy, error)
	SetCriteria(id string, input []model.Criteria) (*model.Hierarchy, error)
}

func NewHierarchyCRUD(repo repository.HierarchyRepository, service *service.CriteriaService) *hierarchyCRUDImpl {
	return &hierarchyCRUDImpl{
		repo:    repo,
		service: service,
	}
}

type hierarchyCRUDImpl struct {
	repo    repository.HierarchyRepository
	service *service.CriteriaService
}

func (hCRUD hierarchyCRUDImpl) RegisterHierarchy(name string, description string, owner string, alternatives []string, objective string) (*model.Hierarchy, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	h := model.NewHierarchy(strings.ReplaceAll(uid.String(), "-", ""), name, description, owner, alternatives, objective)
	errInsert := hCRUD.repo.Save(h)

	return h, errInsert
}

func (hCRUD hierarchyCRUDImpl) GetHierarchy(id string) (*model.Hierarchy, error) {
	return hCRUD.repo.Get(id)
}

func (hCRUD hierarchyCRUDImpl) SetCriteria(id string, input []model.Criteria) (*model.Hierarchy, error) {
	h, err := hCRUD.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if h == nil {
		return nil, fmt.Errorf("the hierarchy %s must be created before set criteria", id)
	}

	h.Criteria = input
	return h, hCRUD.repo.Save(h)
}
