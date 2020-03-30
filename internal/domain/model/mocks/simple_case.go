package mocks

import "dast-api/internal/domain/model"

type mockCase struct {
	H model.Hierarchy
	J model.CriteriaJudgements
}

var SimpleCaseOneLevelCriteria = mockCase{
	H: model.Hierarchy{
		ID:          "123",
		Name:        "Elegír auto",
		Description: "Esta es una jerarquía de ejemplo para elegir un auto",
		Owner:       "amarini",
		Alternatives: []string{
			"fiat",
			"ford",
			"nissan",
		},
		Criteria: []model.Criteria{
			{Level: 0, ID: "velocidad", Name: "Velocidad", Parent: "", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
			{Level: 0, ID: "aceleracion", Name: "Aceleracion", Parent: "", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
			{Level: 0, ID: "cilindrada", Name: "Cilindrada", Parent: "", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
		},
	},
	J: model.CriteriaJudgements{
		ID:          "321",
		HierarchyID: "123",
		CriteriaComparison: []model.CriteriaPairwiseComparison{
			{Level: 0, MatrixContext: struct {
				ComparedTo string
				Elements   []string
				Judgements [][]float64
			}{ComparedTo: "", Elements: []string{"velocidad", "aceleracion", "cilindrada"}, Judgements: [][]float64{
				{
					1.0,
					2.0,
					0.5,
				},
				{
					0.5,
					1.0,
					3.0,
				},
				{
					2.0,
					0.333333333333333,
					1,
				},
			}}},
		},
		AlternativeComparison: []model.MatrixContext{
			{ComparedTo: "velocidad", Elements: []string{"fiat", "ford", "nissan"}, Judgements: [][]float64{
				{
					1.0,
					0.5,
					3,
				},
				{
					2,
					1.0,
					0.5,
				},
				{
					0.333333333333333,
					2,
					1,
				},
			}},
			{ComparedTo: "aceleracion", Elements: []string{"fiat", "ford", "nissan"}, Judgements: [][]float64{
				{
					1.0,
					2,
					3,
				},
				{
					0.5,
					1.0,
					4,
				},
				{
					0.333333333333333,
					0.25,
					1,
				},
			}},
			{ComparedTo: "cilindrada", Elements: []string{"fiat", "ford", "nissan"}, Judgements: [][]float64{
				{
					1.0,
					2,
					0.5,
				},
				{
					0.5,
					1.0,
					0.25,
				},
				{
					2,
					4,
					1,
				},
			}},
		},
		Results: nil,
	},
}
