package mongodb

import (
	"context"
	"dast-api/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoCriteriaJudgements struct {
	ID                    primitive.ObjectID                `bson:"_id"`
	HierarchyID           string                            `bson:"hierarchy_id"`
	Status                string                            `bson:"status"`
	DateCreated           primitive.DateTime                `bson:"date_created"`
	DateLastUpdated       primitive.DateTime                `bson:"date_last_updated"`
	CriteriaComparison    []mongoCriteriaPairwiseComparison `bson:"criteria_comparison"`
	AlternativeComparison []mongoMatrixContext              `bson:"alternative_comparison"`
	Results               map[string]float64                `bson:"results"`
}

func parseToMongoJudgement(j *model.CriteriaJudgements) mongoCriteriaJudgements {
	toRet := mongoCriteriaJudgements{
		ID:                    primitive.NewObjectID(),
		HierarchyID:           j.HierarchyID,
		Status:                j.Status.ToString(),
		DateCreated:           primitive.NewDateTimeFromTime(j.DateCreated),
		DateLastUpdated:       primitive.NewDateTimeFromTime(j.DateLastUpdated),
		CriteriaComparison:    []mongoCriteriaPairwiseComparison{},
		AlternativeComparison: []mongoMatrixContext{},
		Results:               j.Results,
	}

	for _, comparison := range j.CriteriaComparison {
		toRet.CriteriaComparison = append(toRet.CriteriaComparison, mongoCriteriaPairwiseComparison{
			Level: comparison.Level,
			MatrixContext: mongoMatrixContext{
				ComparedTo: comparison.MatrixContext.ComparedTo,
				Elements:   comparison.MatrixContext.Elements,
				Judgements: comparison.MatrixContext.Judgements,
			},
		})
	}

	for _, matrixContext := range j.AlternativeComparison {
		toRet.AlternativeComparison = append(toRet.AlternativeComparison, mongoMatrixContext{
			ComparedTo: matrixContext.ComparedTo,
			Elements:   matrixContext.Elements,
			Judgements: matrixContext.Judgements,
		})
	}

	return toRet
}

func parseMongoMapToNativeJudgement(j mongoCriteriaJudgements) *model.CriteriaJudgements {
	toRet := model.CriteriaJudgements{
		ID:                    j.ID.Hex(),
		HierarchyID:           j.HierarchyID,
		Status:                model.JudgementStatus(j.Status),
		DateCreated:           j.DateCreated.Time(),
		DateLastUpdated:       j.DateLastUpdated.Time(),
		CriteriaComparison:    []model.CriteriaPairwiseComparison{},
		AlternativeComparison: []model.MatrixContext{},
		Results:               j.Results,
	}

	for _, comparison := range j.CriteriaComparison {
		toRet.CriteriaComparison = append(toRet.CriteriaComparison, model.CriteriaPairwiseComparison{
			Level: comparison.Level,
			MatrixContext: model.MatrixContext{
				ComparedTo: comparison.MatrixContext.ComparedTo,
				Elements:   comparison.MatrixContext.Elements,
				Judgements: comparison.MatrixContext.Judgements,
			},
		})
	}

	for _, matrixContext := range j.AlternativeComparison {
		toRet.AlternativeComparison = append(toRet.AlternativeComparison, model.MatrixContext{
			ComparedTo: matrixContext.ComparedTo,
			Elements:   matrixContext.Elements,
			Judgements: matrixContext.Judgements,
		})
	}

	return &toRet
}

func parseMongoMapToNativeJudgements(jj []mongoCriteriaJudgements) []*model.CriteriaJudgements {
	var toRet = []*model.CriteriaJudgements{}
	for _, j := range jj {
		toRet = append(toRet, parseMongoMapToNativeJudgement(j))
	}

	return toRet
}

type mongoCriteriaPairwiseComparison struct {
	Level         int                `bson:"level"`
	MatrixContext mongoMatrixContext `bson:"matrix_context"`
}

type mongoMatrixContext struct {
	ComparedTo string      `bson:"compared_to"`
	Elements   []string    `bson:"elements"`
	Judgements [][]float64 `bson:"judgements"`
}

type CriteriaJudgementsRepository struct {
	collection *mongo.Collection
}

func NewCriteriaJudgementsRepository(collection *mongo.Collection) *CriteriaJudgementsRepository {
	return &CriteriaJudgementsRepository{collection: collection}
}

func (hr CriteriaJudgementsRepository) Override(id string, judgements *model.CriteriaJudgements) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	jID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mongoJudgement := parseToMongoJudgement(judgements)
	mongoJudgement.ID = jID
	if _, err := hr.collection.ReplaceOne(ctx, bson.M{"_id": jID}, mongoJudgement); err != nil {
		return err
	}

	return nil
}

func (hr CriteriaJudgementsRepository) Save(pWise *model.CriteriaJudgements) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoJudgement := parseToMongoJudgement(pWise)
	if _, err := hr.collection.InsertOne(ctx, mongoJudgement); err != nil {
		return err
	}

	pWise.ID = mongoJudgement.ID.Hex()
	return nil
}

func (hr CriteriaJudgementsRepository) Get(id string) (*model.CriteriaJudgements, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filterCursor, err := hr.collection.Find(ctx, bson.M{"_id": oID})
	if err != nil {
		return nil, err
	}

	var episodesFiltered []mongoCriteriaJudgements
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	return parseMongoMapToNativeJudgement(episodesFiltered[0]), nil
}

func (hr CriteriaJudgementsRepository) SearchByHierarchyId(id string) ([]*model.CriteriaJudgements, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterCursor, err := hr.collection.Find(ctx, bson.M{"hierarchy_id": id})
	if err != nil {
		return nil, err
	}

	var episodesFiltered []mongoCriteriaJudgements
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	return parseMongoMapToNativeJudgements(episodesFiltered), nil
}
