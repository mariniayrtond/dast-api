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

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name: "hierarchy-usecase",
			Build: buildHierarchyUsecase,
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
	repo := memory.NewHierarchyRepository()
	service := service.NewCriteriaService(repo)
	return usecase.NewHierarchyCRUD(repo, service), nil

}
