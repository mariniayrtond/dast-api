package model

import "errors"

type CriteriaRoot struct {
	Description string
	Criteria    []*CriteriaNode
}

func (root *CriteriaRoot) GetGlobalScorePriorityVector() []float64 {
	ret := []float64{}
	for _, node := range root.Criteria {
		searchLeafNodesAndFillVector(node, &ret)
	}

	return ret
}

func searchLeafNodesAndFillVector(n *CriteriaNode, ret *[]float64) {
	if n.SubCriteria == nil || len(n.SubCriteria) == 0 {
		*ret = append(*ret, n.GlobalScore)
		return
	}
	for _, subCriterion := range n.SubCriteria {
		searchLeafNodesAndFillVector(subCriterion, ret)
	}
}

func (root *CriteriaRoot) SetCriteria(criteria []Criteria) {
	for _, c := range criteria {
		if c.Level == 0 {
			root.Criteria = append(root.Criteria, &CriteriaNode{
				ID:          c.ID,
				Name:        c.Name,
				LocalScore:  c.Score.Local,
				GlobalScore: c.Score.Global,
				SubCriteria: nil,
			})
		}
	}
}

func (root *CriteriaRoot) SetSubCriteria(criteria []Criteria) {
	for _, criterion := range root.Criteria {
		fillSubCriteria(criterion, criterion.ID, criteria)
	}
}

func fillSubCriteria(node *CriteriaNode, parentID string, criteria []Criteria) {
	for _, criterion := range criteria {
		if criterion.Parent == parentID {
			if node.SubCriteria == nil {
				node.SubCriteria = []*CriteriaNode{}
			}
			node.SubCriteria = append(node.SubCriteria, &CriteriaNode{
				ID:          criterion.ID,
				Name:        criterion.Name,
				LocalScore:  criterion.Score.Local,
				GlobalScore: criterion.Score.Global,
				SubCriteria: nil,
			})
		}
	}

	if node.SubCriteria == nil {
		return
	}

	for _, criterion := range node.SubCriteria {
		fillSubCriteria(criterion, criterion.ID, criteria)
	}
}

func (root *CriteriaRoot) SetScores(j *CriteriaJudgements) error {
	for _, criteriaPairwiseComparison := range j.CriteriaComparison {
		for _, judgement := range criteriaPairwiseComparison.MatrixContext.Judgements {
			for _, f := range judgement {
				if f == 0.0 {
					return errors.New("criteria judgements are not set")
				}
			}
		}

		ahpMatrix := NewAHPMatrix(criteriaPairwiseComparison.MatrixContext.Judgements)
		ahpMatrix.Normalize()
		priorityVector := ahpMatrix.GetPriorityVector()
		for i, alternative := range criteriaPairwiseComparison.MatrixContext.Elements {
			seekAndSetLocalCriteria(root.Criteria, alternative, priorityVector[i])
		}
	}

	root.setGlobalCriteria(root.Criteria)

	return nil
}

func (root *CriteriaRoot) setGlobalCriteria(nodes []*CriteriaNode) {
	for _, node := range nodes {
		node.GlobalScore = node.LocalScore
		if node.SubCriteria != nil && len(node.SubCriteria) > 0 {
			for _, subNode := range node.SubCriteria {
				subNode.GlobalScore = subNode.LocalScore * node.GlobalScore
				setGlobalCriteriaForSubCriteria(subNode)
			}
		}
	}
}

func setGlobalCriteriaForSubCriteria(node *CriteriaNode) {
	if node.SubCriteria != nil && len(node.SubCriteria) > 0 {
		for _, subNode := range node.SubCriteria {
			subNode.GlobalScore = subNode.LocalScore * node.GlobalScore
			setGlobalCriteriaForSubCriteria(subNode)
		}
	}
}

func seekAndSetLocalCriteria(nodes []*CriteriaNode, name string, val float64) {
	for _, node := range nodes {
		if node.ID == name {
			node.LocalScore = val
			return
		}

		if node.SubCriteria != nil && len(node.SubCriteria) > 0 {
			seekAndSetLocalCriteria(node.SubCriteria, name, val)
		}
	}
}

type CriteriaNode struct {
	ID          string
	Name        string
	LocalScore  float64
	GlobalScore float64
	SubCriteria []*CriteriaNode
}

func NewCriteriaHierarchy(description string, criteria []Criteria) CriteriaRoot {
	root := CriteriaRoot{
		Description: description,
		Criteria:    nil,
	}

	root.SetCriteria(criteria)
	root.SetSubCriteria(criteria)

	return root
}
