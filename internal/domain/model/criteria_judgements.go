package model

import (
	"errors"
)

type CriteriaJudgements struct {
	ID                    string
	HierarchyID           string
	CriteriaComparison    []CriteriaPairwiseComparison
	AlternativeComparison []MatrixContext
	Results               map[string]float64
}

func NewCriteriaJudgements(ID string, hierarchyID string, criteriaPairwiseComparisons []CriteriaPairwiseComparison, alternativeComparison []MatrixContext) *CriteriaJudgements {
	return &CriteriaJudgements{ID: ID, HierarchyID: hierarchyID, CriteriaComparison: criteriaPairwiseComparisons, AlternativeComparison: alternativeComparison}
}

func (j *CriteriaJudgements) GenerateResults(tree *CriteriaRoot, alternatives []string) error {
	alternativePriorityMatrix := [][]float64{}
	for _, alternativeMatrixContext := range j.AlternativeComparison {
		for _, judgement := range alternativeMatrixContext.Judgements {
			for _, f := range judgement {
				if f == 0.0 {
					return errors.New("alternatives judgements are not set")
				}
			}
		}
		ahpMatrix := NewAHPMatrix(alternativeMatrixContext.Judgements)
		ahpMatrix.Normalize()
		alternativePriorityMatrix = append(alternativePriorityMatrix, ahpMatrix.GetPriorityVector())
	}

	criteriaPriorityMatrix := [][]float64{}
	criteriaPriorityMatrix = append(criteriaPriorityMatrix, tree.GetGlobalScorePriorityVector())

	alternativeAHPMatrix := NewAHPMatrix(alternativePriorityMatrix)
	criteriaAHPMatrix := NewAHPMatrix(criteriaPriorityMatrix)

	var err error
	alternativeAHPMatrix.m, err = alternativeAHPMatrix.m.Transpose()
	if err != nil {
		return err
	}
	criteriaAHPMatrix.m, err = criteriaAHPMatrix.m.Transpose()
	if err != nil {
		return err
	}

	productMatrix, err := alternativeAHPMatrix.m.DotProduct(criteriaAHPMatrix.m)
	if err != nil {
		return err
	}

	j.Results = map[string]float64{}
	for i, alternative := range alternatives {
		j.Results[alternative] = productMatrix.At(i, 0)
	}

	return nil
}

type CriteriaPairwiseComparison struct {
	Level         int
	MatrixContext MatrixContext
}

type MatrixContext struct {
	ComparedTo string
	Elements   []string
	Judgements [][]float64
}
