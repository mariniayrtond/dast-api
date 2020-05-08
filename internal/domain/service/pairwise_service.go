package service

import "dast-api/internal/domain/model"

type PairwiseService struct{}

func NewPairwiseService() *PairwiseService {
	return &PairwiseService{}
}

func (s PairwiseService) GenerateCriteriaMatrices(root *model.CriteriaRoot) []model.CriteriaPairwiseComparison {
	ret := []model.CriteriaPairwiseComparison{}
	generateCriteriaMatrixByTree(&ret, 0, root.Criteria, nil)
	return ret
}

func generateCriteriaMatrixByTree(res *[]model.CriteriaPairwiseComparison, level int, criteriaRange []*model.CriteriaNode, caller *model.CriteriaNode) {
	if criteriaRange == nil {
		return
	}

	comparedTo := ""
	if caller != nil {
		comparedTo = caller.ID
	}

	comparison := model.CriteriaPairwiseComparison{
		Level: level,
		MatrixContext: model.MatrixContext{
			ComparedTo: comparedTo,
			Elements:   []string{},
			Judgements: [][]float64{},
		},
	}

	for i, node := range criteriaRange {
		comparison.MatrixContext.Elements = append(comparison.MatrixContext.Elements, node.ID)
		comparison.MatrixContext.Judgements = append(comparison.MatrixContext.Judgements, []float64{})

		for k := 0; k < len(criteriaRange); k++ {
			if k == i {
				comparison.MatrixContext.Judgements[i] = append(comparison.MatrixContext.Judgements[i], 1.0)
			} else {
				comparison.MatrixContext.Judgements[i] = append(comparison.MatrixContext.Judgements[i], 0.0)
			}

		}
	}

	*res = append(*res, comparison)
	for _, node := range criteriaRange {
		generateCriteriaMatrixByTree(res, level+1, node.SubCriteria, node)
	}
}

func (s PairwiseService) GenerateAlternativeMatrices(root *model.CriteriaRoot, alternatives []string) []model.MatrixContext {
	ret := []model.MatrixContext{}
	for _, criterion := range root.Criteria {
		generateAlternativeMatrixByTree(&ret, criterion, alternatives)
	}
	return ret
}

func generateAlternativeMatrixByTree(res *[]model.MatrixContext, caller *model.CriteriaNode, alternatives []string) {
	if caller.SubCriteria == nil || len(caller.SubCriteria) == 0 {
		judgements := [][]float64{}
		for k := 0; k < len(alternatives); k++ {
			judgements = append(judgements, []float64{})
			for j := 0; j < len(alternatives); j++ {
				if k == j {
					judgements[k] = append(judgements[k], 1.0)
				} else {
					judgements[k] = append(judgements[k], 0.0)
				}
			}
		}

		*res = append(*res, model.MatrixContext{
			ComparedTo: caller.ID,
			Elements:   alternatives,
			Judgements: judgements,
		})
		return
	}

	for _, criterion := range caller.SubCriteria {
		generateAlternativeMatrixByTree(res, criterion, alternatives)
	}


}
