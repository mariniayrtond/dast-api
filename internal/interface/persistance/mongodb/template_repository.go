package mongodb

import (
	"context"
	"dast-api/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoCriteriaTemplate struct {
	ID          primitive.ObjectID `bson:"_id"`
	Owner       string             `bson:"owner"`
	Description string             `bson:"description"`
	Criteria    []mongoCriteria    `bson:"criteria"`
}

func parseToMongoCriteriaTemplate(t *model.CriteriaTemplate) mongoCriteriaTemplate {
	return mongoCriteriaTemplate{
		ID:          primitive.NewObjectID(),
		Owner:       t.Owner,
		Description: t.Description,
		Criteria:    parseToMongoCriteria(t.Criteria),
	}
}

func parseMongoMapToNativeCriteriaTemplate(t mongoCriteriaTemplate) *model.CriteriaTemplate {
	return &model.CriteriaTemplate{
		ID:          t.ID.Hex(),
		Owner:       t.Owner,
		Description: t.Description,
		Criteria:    parseMongoMapToNativeCriteria(t.Criteria),
	}
}

func parseMongoMapToNativeCriteriaTemplates(ts []mongoCriteriaTemplate) []*model.CriteriaTemplate {
	toRet := []*model.CriteriaTemplate{}
	for _, template := range ts {
		toRet = append(toRet, parseMongoMapToNativeCriteriaTemplate(template))
	}
	return toRet
}

type TemplateRepository struct {
	collection *mongo.Collection
}

func NewTemplateRepository(collection *mongo.Collection) *TemplateRepository {
	return &TemplateRepository{collection: collection}
}

func (t TemplateRepository) Save(template *model.CriteriaTemplate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoTemplate := parseToMongoCriteriaTemplate(template)
	if _, err := t.collection.InsertOne(ctx, mongoTemplate); err != nil {
		return err
	}

	template.ID = mongoTemplate.ID.Hex()
	return nil
}

func (t TemplateRepository) SearchPublicTemplates() ([]*model.CriteriaTemplate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterCursor, err := t.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var episodesFiltered []mongoCriteriaTemplate
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	return parseMongoMapToNativeCriteriaTemplates(episodesFiltered), nil
}
