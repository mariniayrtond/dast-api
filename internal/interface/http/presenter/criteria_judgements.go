package presenter

import "dast-api/internal/domain/model"

type criteriaJudgements struct {
	ID                    string               `json:"id"`
	HierarchyID           string               `json:"hierarchy_id"`
	CriteriaComparison    []pairwiseComparison `json:"criteria_comparison"`
	AlternativeComparison []matrixContext      `json:"alternative_comparison"`
	Results               map[string]float64   `json:"results"`
}

type pairwiseComparison struct {
	Level         int           `json:"level"`
	MatrixContext matrixContext `json:"matrix_context"`
}

type matrixContext struct {
	ComparedTo string      `json:"compared_to"`
	Elements   []string    `json:"elements"`
	Judgements [][]float64 `json:"judgements"`
}

func RenderCriteriaJudgements(j *model.CriteriaJudgements) criteriaJudgements {
	ret := criteriaJudgements{
		ID:                 j.ID,
		HierarchyID:        j.HierarchyID,
		CriteriaComparison: []pairwiseComparison{},
	}

	for _, comparison := range j.CriteriaComparison {
		ret.CriteriaComparison = append(ret.CriteriaComparison, pairwiseComparison{
			Level: comparison.Level,
			MatrixContext: matrixContext{
				ComparedTo: comparison.MatrixContext.ComparedTo,
				Elements:   comparison.MatrixContext.Elements,
				Judgements: comparison.MatrixContext.Judgements,
			},
		})
	}

	for _, alternativeC := range j.AlternativeComparison {
		ret.AlternativeComparison = append(ret.AlternativeComparison, matrixContext{
			ComparedTo: alternativeC.ComparedTo,
			Elements:   alternativeC.Elements,
			Judgements: alternativeC.Judgements,
		})
	}

	if j.Results != nil {
		ret.Results = j.Results
	}

	return ret
}
