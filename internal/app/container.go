package app

import (
	"context"
	"dast-api/internal/config"
	"dast-api/internal/domain/service"
	"dast-api/internal/interface/persistance/mongodb"
	"dast-api/internal/usecase"
	"fmt"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

type Container struct {
	ctn di.Container
}

var hRepo *mongodb.HierarchyRepository
var uRepo *mongodb.UserRepository
var pRepo *mongodb.CriteriaJudgementsRepository
var tRepo *mongodb.TemplateRepository
var tokenRepo *mongodb.TokenRepository

func NewContainer() (*Container, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf(os.Getenv("mongo"), config.MongoDBName),
	))

	if err != nil {
		return nil, err
	}

	if errPing := client.Ping(ctx, readpref.Primary()); errPing != nil {
		return nil, errPing
	}

	hRepo = mongodb.NewHierarchyRepository(client.Database(config.MongoDBName).Collection(config.MongoTableHierarchies))
	uRepo = mongodb.NewUserRepository(client.Database(config.MongoDBName).Collection(config.MongoTableUsers))
	pRepo = mongodb.NewCriteriaJudgementsRepository(client.Database(config.MongoDBName).Collection(config.MongoTableJudgements))
	tRepo = mongodb.NewTemplateRepository(client.Database(config.MongoDBName).Collection(config.MongoTableTemplates))
	tokenRepo = mongodb.NewTokenRepository(client.Database(config.MongoDBName).Collection(config.MongoTableTokens))

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
	criteriaService := service.NewCriteriaService(hRepo)
	return usecase.NewHierarchyCRUD(hRepo, tRepo, criteriaService), nil
}

func buildPairwiseUsecase(ctn di.Container) (interface{}, error) {
	pairwiseService := service.NewPairwiseService()
	return usecase.NewPairwiseComparisonUC(hRepo, pRepo, pairwiseService), nil
}

func buildUserUsecase(ctn di.Container) (interface{}, error) {
	return usecase.NewUserUseCase(uRepo, tokenRepo), nil
}
