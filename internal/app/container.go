package app

import (
	"dast-api/internal/domain/service"
	"dast-api/internal/interface/persistance/memory"
	"dast-api/internal/usecase"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

var hRepo *memory.HierarchyRepository
var pRepo *memory.CriteriaJudgementsRepository

func NewContainer() (*Container, error) {
	hRepo = memory.NewHierarchyRepository()
	pRepo = memory.NewCriteriaJudgementsRepository()

	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "hierarchy-usecase",
			Build: buildHierarchyUsecase,
		},
		{
			Name:  "pwise-usecase",
			Build: buildPairwiseUsecase,
		},
		{
			Name:  "user-usecase",
			Build: buildUserUsecase,
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildHierarchyUsecase(ctn di.Container) (interface{}, error) {
	service := service.NewCriteriaService(hRepo)
	return usecase.NewHierarchyCRUD(hRepo, service), nil
}

func buildPairwiseUsecase(ctn di.Container) (interface{}, error) {
	service := service.NewPairwiseService()
	return usecase.NewPairwiseComparisonUC(hRepo, pRepo, service), nil
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	return usecase.NewUserUseCase(memory.NewUserRepository(), memory.NewTokenRepository()), nil
}
