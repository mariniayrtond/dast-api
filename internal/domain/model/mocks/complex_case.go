package mocks

import "dast-api/internal/domain/model"

var ComplexCaseTwoLevelCriteria = mockCase{
	H: model.Hierarchy{
		ID:          "123",
		Name:        "chooseCar",
		Description: "Esta es una jerarquía de ejemplo para elegir un auto",
		Owner:       "amarini",
		Alternatives: []string{
			"ford",
			"fiat",
			"chevrolet",
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
			{Level: 1, ID: "velocidad1", Name: "Velocidad1", Parent: "velocidad", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
			{Level: 1, ID: "velocidad2", Name: "Velocidad2", Parent: "velocidad", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
			{Level: 1, ID: "aceleracion1", Name: "Aceleracion1", Parent: "aceleracion", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
			{Level: 1, ID: "aceleracion2", Name: "Aceleracion2", Parent: "aceleracion", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
			{Level: 2, ID: "aceleracion21", Name: "Aceleracion21", Parent: "aceleracion2", Score: struct {
				Local  float64
				Global float64
			}{Local: 0.0, Global: 0.0}},
			{Level: 2, ID: "aceleracion22", Name: "Aceleracion22", Parent: "aceleracion2", Score: struct {
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
			}{ComparedTo: "", Elements: []string{"velocidad", "cilindrada", "aceleracion"}, Judgements: [][]float64{
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
			{Level: 1, MatrixContext: struct {
				ComparedTo string
				Elements   []string
				Judgements [][]float64
			}{ComparedTo: "velocidad", Elements: []string{"velocidad1", "velocidad2"}, Judgements: [][]float64{
				{
					1.0,
					2.0,
				},
				{
					0.5,
					1.0,
				},
			}}},
			{Level: 1, MatrixContext: struct {
				ComparedTo string
				Elements   []string
				Judgements [][]float64
			}{ComparedTo: "aceleracion", Elements: []string{"aceleracion1", "aceleracion2"}, Judgements: [][]float64{
				{
					1.0,
					5.0,
				},
				{
					0.2,
					1.0,
				},
			}}},
			{Level: 2, MatrixContext: struct {
				ComparedTo string
				Elements   []string
				Judgements [][]float64
			}{ComparedTo: "aceleracion2", Elements: []string{"aceleracion21", "aceleracion22"}, Judgements: [][]float64{
				{
					1.0,
					0.4,
				},
				{
					2.5,
					1.0,
				},
			}}},
		},
		AlternativeComparison: []model.MatrixContext{
			{ComparedTo: "velocidad1", Elements: []string{"ford", "fiat", "chevrolet"}, Judgements: [][]float64{
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
			{ComparedTo: "velocidad2", Elements: []string{"ford", "fiat", "chevrolet"}, Judgements: [][]float64{
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
			{ComparedTo: "aceleracion1", Elements: []string{"ford", "fiat", "chevrolet"}, Judgements: [][]float64{
				{
					1.0,
					2,
					0.5,
				},
				{
					0.5,
					1.0,
					0.333333333333333,
				},
				{
					2,
					3,
					1,
				},
			}},
			{ComparedTo: "aceleracion21", Elements: []string{"ford", "fiat", "chevrolet"}, Judgements: [][]float64{
				{
					1.0,
					4,
					5,
				},
				{
					0.25,
					1.0,
					3,
				},
				{
					0.2,
					0.333333333333333,
					1,
				},
			}},
			{ComparedTo: "aceleracion22", Elements: []string{"ford", "fiat", "chevrolet"}, Judgements: [][]float64{
				{
					1.0,
					2,
					1,
				},
				{
					0.5,
					1.0,
					0.333333333333333,
				},
				{
					1,
					3,
					1,
				},
			}},
			{ComparedTo: "cilindrada", Elements: []string{"ford", "fiat", "chevrolet"}, Judgements: [][]float64{
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
