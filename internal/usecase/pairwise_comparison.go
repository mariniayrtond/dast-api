package usecase

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/domain/repository"
	"dast-api/internal/domain/service"
	"dast-api/pkg/uid"
	"fmt"
)

type PairwiseComparison interface {
	GenerateMatrices(hierarchy *model.Hierarchy) (*model.CriteriaJudgements, error)
	GetJudgements(idHierarchy string, idJudgements string) (*model.CriteriaJudgements, error)
	UpdateJudgements(idHierarchy string, idJudgements string, judgements *model.CriteriaJudgements) (*model.CriteriaJudgements, error)
	GenerateResults(hierarchy *model.Hierarchy, idJudgements string) (*model.CriteriaJudgements, error)
}

func NewPairwiseComparisonUC(hRepo repository.HierarchyRepository, pRepo repository.CriteriaJudgementsRepository, service *service.PairwiseService) PairwiseComparison {
	return pairwiseComparisonImpl{
		hRepo:   hRepo,
		pRepo:   pRepo,
		service: service,
	}
}

type pairwiseComparisonImpl struct {
	hRepo   repository.HierarchyRepository
	pRepo   repository.CriteriaJudgementsRepository
	service *service.PairwiseService
}

func (p pairwiseComparisonImpl) GenerateResults(h *model.Hierarchy, idJudgements string) (*model.CriteriaJudgements, error) {
	j, err := p.pRepo.Get(idJudgements)
	if err != nil {
		return nil, err
	}

	if j == nil {
		return nil, fmt.Errorf("judgements:%s not found", idJudgements)
	}

	if j.HierarchyID != h.ID {
		return nil, fmt.Errorf("judgements_id:%s does not belong to the hierarchy:%s", h.ID, idJudgements)
	}

	tree := model.NewCriteriaHierarchy(h.Description, h.Criteria)

	if err := tree.SetScores(j); err != nil {
		return nil, err
	}

	if err := j.GenerateResults(&tree, h.Alternatives); err != nil {
		return nil, err
	}

	if err := p.pRepo.Override(j.ID, j); err != nil {
		return nil, err
	}

	return j, nil
}

func (p pairwiseComparisonImpl) GetJudgements(idHierarchy string, idJudgments string) (*model.CriteriaJudgements, error) {
	j, err := p.pRepo.Get(idJudgments)
	if err != nil {
		return nil, err
	}

	if j == nil {
		return nil, fmt.Errorf("judgements:%s doesnot exists", idJudgments)
	}

	if j.HierarchyID != idHierarchy {
		return nil, fmt.Errorf("judgements:%s does not append to hierarcht:%s", idJudgments, idHierarchy)
	}

	return j, nil
}

func (p pairwiseComparisonImpl) UpdateJudgements(idHierarchy string, idJudgments string, judgements *model.CriteriaJudgements) (*model.CriteriaJudgements, error) {
	j, err := p.pRepo.Get(idJudgments)
	if err != nil {
		return nil, err
	}

	if j == nil {
		return nil, fmt.Errorf("judgements:%s doesnot exists", idJudgments)
	}

	if j.HierarchyID != idHierarchy {
		return nil, fmt.Errorf("judgements:%s does not append to hierarcht:%s", idJudgments, idHierarchy)
	}

	if err := p.service.CheckCriteriaOrder(j, judgements); err != nil {
		return nil, err
	}

	judgements.HierarchyID = j.HierarchyID
	judgements.ID = j.ID
	judgements.Results = j.Results

	if err := p.pRepo.Override(j.ID, judgements); err != nil {
		return nil, err
	}

	return judgements, nil
}

func (p pairwiseComparisonImpl) GenerateMatrices(h *model.Hierarchy) (*model.CriteriaJudgements, error) {
	tree := model.NewCriteriaHierarchy(h.Description, h.Criteria)

	jID, err := uid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	judgments := model.NewCriteriaJudgements(jID, h.ID, p.service.GenerateCriteriaMatrices(&tree), p.service.GenerateAlternativeMatrices(&tree, h.Alternatives))
	if err := p.pRepo.Save(judgments); err != nil {
		return nil, err
	}

	return judgments, nil
}
