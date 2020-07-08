package presenter

import (
	"dast-api/internal/domain/model"
	"time"
)

type CriteriaJudgements struct {
	ID                    string               `json:"id"`
	HierarchyID           string               `json:"hierarchy_id"`
	Status                string               `json:"status"`
	DateCreated           string               `json:"date_created"`
	DateLastUpdated       string               `json:"date_last_updated"`
	CriteriaComparison    []PairwiseComparison `json:"criteria_comparison"`
	AlternativeComparison []MatrixContext      `json:"alternative_comparison"`
	Results               map[string]float64   `json:"results,omitempty"`
}

type PairwiseComparison struct {
	Level         int           `json:"level"`
	MatrixContext MatrixContext `json:"matrix_context"`
}

type MatrixContext struct {
	ComparedTo string      `json:"compared_to"`
	Elements   []string    `json:"elements"`
	Judgements [][]float64 `json:"judgements"`
}

func RenderSomeCriteriaJudgements(jj []*model.CriteriaJudgements) []CriteriaJudgements {
	var ret []CriteriaJudgements
	for _, j := range jj {
		ret = append(ret, RenderCriteriaJudgements(j))
	}

	return ret
}

func RenderCriteriaJudgements(j *model.CriteriaJudgements) CriteriaJudgements {
	ret := CriteriaJudgements{
		ID:                 j.ID,
		Status:             j.Status.ToString(),
		DateLastUpdated:    j.DateLastUpdated.Format(time.RFC3339),
		DateCreated:        j.DateCreated.Format(time.RFC3339),
		HierarchyID:        j.HierarchyID,
		CriteriaComparison: []PairwiseComparison{},
	}

	for _, comparison := range j.CriteriaComparison {
		ret.CriteriaComparison = append(ret.CriteriaComparison, PairwiseComparison{
			Level: comparison.Level,
			MatrixContext: MatrixContext{
				ComparedTo: comparison.MatrixContext.ComparedTo,
				Elements:   comparison.MatrixContext.Elements,
				Judgements: comparison.MatrixContext.Judgements,
			},
		})
	}

	for _, alternativeC := range j.AlternativeComparison {
		ret.AlternativeComparison = append(ret.AlternativeComparison, MatrixContext{
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
