package repository

import "dast-api/internal/domain/model"

//go:generate moq -out criteria_judgements_repository_moq.go . CriteriaJudgementsRepository
type CriteriaJudgementsRepository interface {
	Override(id string, judgements *model.CriteriaJudgements) error
	Save(judgements *model.CriteriaJudgements) error
	Get(id string) (*model.CriteriaJudgements, error)
}
