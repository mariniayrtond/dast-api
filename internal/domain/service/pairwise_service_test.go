package service

import (
	"dast-api/internal/domain/model"
	"reflect"
	"testing"
)



func TestPairwiseService_GenerateCriteriaMatrices1(t *testing.T) {
	type args struct {
		root *model.CriteriaRoot
	}
	tests := []struct {
		name string
		args args
		want []model.CriteriaPairwiseComparison
	}{
		{
			name: "test-happy-1",
			args: args{
				root: &model.CriteriaRoot{
					Description: "Elegí el mejor plan",
					Criteria: []*model.CriteriaNode{
						&model.CriteriaNode{
							ID:          "velocity",
							Name:        "velocidad",
							LocalScore:  0,
							GlobalScore: 0,
							SubCriteria: []*model.CriteriaNode{
								&model.CriteriaNode{
									ID:          "velocity11",
									Name:        "velocity11",
									LocalScore:  0,
									GlobalScore: 0,
									SubCriteria: nil,
								},
								&model.CriteriaNode{
									ID:          "velocity12",
									Name:        "velocity12",
									LocalScore:  0,
									GlobalScore: 0,
									SubCriteria: nil,
								},
							},
						},
						&model.CriteriaNode{
							ID:          "aceleration",
							Name:        "aceleracion",
							LocalScore:  0,
							GlobalScore: 0,
							SubCriteria: []*model.CriteriaNode{
								&model.CriteriaNode{
									ID:          "aceleration11",
									Name:        "aceleration11",
									LocalScore:  0,
									GlobalScore: 0,
									SubCriteria: nil,
								},
								&model.CriteriaNode{
									ID:          "aceleration12",
									Name:        "aceleration12",
									LocalScore:  0,
									GlobalScore: 0,
									SubCriteria: []*model.CriteriaNode{
										&model.CriteriaNode{
											ID:          "aceleration121",
											Name:        "aceleration121",
											LocalScore:  0,
											GlobalScore: 0,
											SubCriteria: nil,
										},
										&model.CriteriaNode{
											ID:          "aceleration122",
											Name:        "aceleration122",
											LocalScore:  0,
											GlobalScore: 0,
											SubCriteria: nil,
										},
									},
								},
							},
						},
						&model.CriteriaNode{
							ID:          "comodidad",
							Name:        "comodidad",
							LocalScore:  0,
							GlobalScore: 0,
							SubCriteria: nil,
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PairwiseService{}
			if got := s.GenerateCriteriaMatrices(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateMatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPairwiseService_GenerateAlternativeMatrices(t *testing.T) {
	type args struct {
		root         *model.CriteriaRoot
		alternatives []string
	}
	tests := []struct {
		name string
		args args
		want []model.MatrixContext
	}{
		{
			name: "happy-case-1",
			args: args{
				root:         &model.CriteriaRoot{
					Description: "elegi el mejor auto",
					Criteria: []*model.CriteriaNode{
						{
							ID:          "velocity",
							Name:        "velocity",
							LocalScore:  0,
							GlobalScore: 0,
							SubCriteria: nil,
						},
						{
							ID:          "celerity",
							Name:        "celerity",
							LocalScore:  0,
							GlobalScore: 0,
							SubCriteria: []*model.CriteriaNode{
								{
									ID:          "celerity11",
									Name:        "celerity11",
									LocalScore:  0,
									GlobalScore: 0,
									SubCriteria: nil,
								},
								{
									ID:          "celerity12",
									Name:        "celerity12",
									LocalScore:  0,
									GlobalScore: 0,
									SubCriteria: nil,
								},
							},
						},
					},
				},
				alternatives: []string{
					"auto1",
					"auto2",
					"auto3",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := PairwiseService{}
			if got := s.GenerateAlternativeMatrices(tt.args.root, tt.args.alternatives); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateAlternativeMatrices() = %v, want %v", got, tt.want)
			}
		})
	}
}