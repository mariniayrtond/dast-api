package usecase

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/domain/repository"
	"dast-api/internal/domain/service"
	"dast-api/pkg/uid"
)

type HierarchyCRUD interface {
	RegisterHierarchy(string, string, string, []string, string) (*model.Hierarchy, error)
	GetHierarchy(id string) (*model.Hierarchy, error)
	SetCriteria(hierarchy *model.Hierarchy, input []model.Criteria) (*model.Hierarchy, error)
	SaveCriteriaTemplate(owner string, description string, criteria []model.Criteria) (*model.CriteriaTemplate, error)
	SearchPublicTemplates() ([]*model.CriteriaTemplate, error)
	SearchByUsername(username string) ([]*model.Hierarchy, error)
}

func NewHierarchyCRUD(repo repository.HierarchyRepository, templatesRepo repository.TemplateRepository, service *service.CriteriaService) HierarchyCRUD {
	return hierarchyCRUDImpl{
		repo:          repo,
		templatesRepo: templatesRepo,
		service:       service,
	}
}

type hierarchyCRUDImpl struct {
	repo          repository.HierarchyRepository
	templatesRepo repository.TemplateRepository
	service       *service.CriteriaService
}

func (hCRUD hierarchyCRUDImpl) SaveCriteriaTemplate(owner string, description string, criteria []model.Criteria) (*model.CriteriaTemplate, error) {
	id, err := uid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	template := model.NewCriteriaTemplate(id, description, owner, criteria)
	errInsert := hCRUD.templatesRepo.Save(template)
	return template, errInsert
}

func (hCRUD hierarchyCRUDImpl) SearchPublicTemplates() ([]*model.CriteriaTemplate, error) {
	return hCRUD.templatesRepo.SearchPublicTemplates()
}

func (hCRUD hierarchyCRUDImpl) RegisterHierarchy(name string, description string, owner string, alternatives []string, objective string) (*model.Hierarchy, error) {
	id, err := uid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	h := model.NewHierarchy(id, name, description, owner, alternatives)
	errInsert := hCRUD.repo.Save(h)

	return h, errInsert
}

func (hCRUD hierarchyCRUDImpl) GetHierarchy(id string) (*model.Hierarchy, error) {
	return hCRUD.repo.Get(id)
}

func (hCRUD hierarchyCRUDImpl) SetCriteria(h *model.Hierarchy, input []model.Criteria) (*model.Hierarchy, error) {
	h.Criteria = input
	return h, hCRUD.repo.Override(h.ID, h)
}

func (hCRUD hierarchyCRUDImpl) SearchByUsername(username string) ([]*model.Hierarchy, error) {
	return hCRUD.repo.SearchByUsername(username)
}
