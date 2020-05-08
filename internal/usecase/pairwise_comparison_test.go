package usecase

import (
	"dast-api/internal/domain/model"
	"dast-api/internal/domain/model/mocks"
	"dast-api/internal/domain/repository"
	"dast-api/internal/domain/service"
	"errors"
	"reflect"
	"testing"
)

func Test_pairwiseComparisonImpl_GenerateResults(t *testing.T) {
	type fields struct {
		hRepo   repository.HierarchyRepository
		pRepo   repository.CriteriaJudgementsRepository
		service *service.PairwiseService
	}
	type args struct {
		idHierarchy  string
		idJudgements string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]float64
		wantErr bool
	}{
		{
			name:    "sad_case_1",
			fields:  fields{
				hRepo:   &repository.HierarchyRepositoryMock{
					GetFunc: func(id string) (hierarchy *model.Hierarchy, err error) {
						return &model.Hierarchy{}, errors.New("error")
					},
				},
				pRepo:   &repository.CriteriaJudgementsRepositoryMock{
					GetFunc: func(id string) (judgements *model.CriteriaJudgements, err error) {
						return nil, errors.New("error")
					},
				},
				service: service.NewPairwiseService(),
			},
			args:    args{
				idHierarchy:  "123",
				idJudgements: "123",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "happy_case_1",
			fields:  fields{
				hRepo:   &repository.HierarchyRepositoryMock{
					GetFunc: func(id string) (hierarchy *model.Hierarchy, err error) {
						return &mocks.SimpleCaseOneLevelCriteria.H, nil
					},
				},
				pRepo:   &repository.CriteriaJudgementsRepositoryMock{
					GetFunc: func(id string) (judgements *model.CriteriaJudgements, err error) {
						return &mocks.SimpleCaseOneLevelCriteria.J, nil
					},
				},
				service: service.NewPairwiseService(),
			},
			args:    args{
				idHierarchy:  "123",
				idJudgements: "123",
			},
			want: map[string]float64{
				"fiat": 0.3973375580915263,
				"ford": 0.2861561464206967,
				"nissan": 0.31650629548777687,
			},
			wantErr: false,
		},
		{
			name:    "happy_case_2",
			fields:  fields{
				hRepo:   &repository.HierarchyRepositoryMock{
					GetFunc: func(id string) (hierarchy *model.Hierarchy, err error) {
						return &mocks.ComplexCaseTwoLevelCriteria.H, nil
					},
				},
				pRepo:   &repository.CriteriaJudgementsRepositoryMock{
					GetFunc: func(id string) (judgements *model.CriteriaJudgements, err error) {
						return &mocks.ComplexCaseTwoLevelCriteria.J, nil
					},
				},
				service: service.NewPairwiseService(),
			},
			args:    args{
				idHierarchy:  "123",
				idJudgements: "123",
			},
			want: map[string]float64{
				"ford": 0.34125622741095235,
				"fiat": 0.2162846496027463,
				"chevrolet": 0.4424591229863012,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := pairwiseComparisonImpl{
				hRepo:   tt.fields.hRepo,
				pRepo:   tt.fields.pRepo,
				service: tt.fields.service,
			}
			got, err := p.GenerateResults(tt.args.idHierarchy, tt.args.idJudgements)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !reflect.DeepEqual(got.Results, tt.want) {
				t.Errorf("GenerateResults() got = %v, want %v", got, tt.want)
			}
		})
	}
}