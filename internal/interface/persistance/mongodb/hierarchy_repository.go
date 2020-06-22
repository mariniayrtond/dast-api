package mongodb

import (
	"context"
	"dast-api/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoHierarchy struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Owner        string             `bson:"owner"`
	Alternatives []string           `bson:"alternatives"`
	Criteria     []mongoCriteria    `bson:"criteria"`
}

func parseToMongoHierarchy(h *model.Hierarchy) mongoHierarchy {
	return mongoHierarchy{
		ID:           primitive.NewObjectID(),
		Name:         h.Name,
		Description:  h.Description,
		Owner:        h.Owner,
		Alternatives: h.Alternatives,
		Criteria:     parseToMongoCriteria(h.Criteria),
	}
}

func parseMongoMapToNativeHierarchies(mh []mongoHierarchy) []*model.Hierarchy {
	toRet := []*model.Hierarchy{}

	for _, mongoHierarchy := range mh {
		toRet = append(toRet, parseMongoMapToNativeHierarchy(mongoHierarchy))
	}

	return toRet
}

func parseMongoMapToNativeHierarchy(mh mongoHierarchy) *model.Hierarchy {
	return &model.Hierarchy{
		ID:           mh.ID.Hex(),
		Name:         mh.Name,
		Description:  mh.Description,
		Owner:        mh.Owner,
		Alternatives: mh.Alternatives,
		Criteria:     parseMongoMapToNativeCriteria(mh.Criteria),
	}
}

type mongoCriteria struct {
	Level  int        `bson:"level"`
	ID     string     `bson:"id"`
	Name   string     `bson:"name"`
	Parent string     `bson:"parent"`
	Score  mongoScore `bson:"score"`
}

type mongoScore struct {
	Local  float64 `bson:"local"`
	Global float64 `bson:"global"`
}

func parseToMongoCriteria(c []model.Criteria) []mongoCriteria {
	toRet := []mongoCriteria{}
	for _, criterion := range c {
		toRet = append(toRet, mongoCriteria{
			Level:  criterion.Level,
			ID:     criterion.ID,
			Name:   criterion.Name,
			Parent: criterion.Parent,
			Score: mongoScore{
				Local:  criterion.Score.Local,
				Global: criterion.Score.Global,
			},
		})
	}

	return toRet
}

func parseMongoMapToNativeCriteria(mc []mongoCriteria) []model.Criteria {
	toRet := []model.Criteria{}
	for _, criterion := range mc {
		toRet = append(toRet, model.Criteria{
			Level:  criterion.Level,
			ID:     criterion.ID,
			Name:   criterion.Name,
			Parent: criterion.Parent,
			Score: model.Score{
				Local:  criterion.Score.Local,
				Global: criterion.Score.Global,
			},
		})
	}

	return toRet
}

type HierarchyRepository struct {
	collection *mongo.Collection
}

func NewHierarchyRepository(collection *mongo.Collection) *HierarchyRepository {
	return &HierarchyRepository{collection: collection}
}

func (hr HierarchyRepository) Override(id string, hierarchy *model.Hierarchy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	mongoHierarchy := parseToMongoHierarchy(hierarchy)
	mongoHierarchy.ID = hID
	if _, err := hr.collection.ReplaceOne(ctx, bson.M{"_id": hID}, mongoHierarchy); err != nil {
		return err
	}

	return nil
}

func (hr HierarchyRepository) Save(hierarchy *model.Hierarchy) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoHierarchy := parseToMongoHierarchy(hierarchy)
	if _, err := hr.collection.InsertOne(ctx, mongoHierarchy); err != nil {
		return err
	}

	hierarchy.ID = mongoHierarchy.ID.Hex()
	return nil
}

func (hr HierarchyRepository) Get(id string) (*model.Hierarchy, error) {
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

	var episodesFiltered []mongoHierarchy
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	return parseMongoMapToNativeHierarchy(episodesFiltered[0]), nil
}

func (hr HierarchyRepository) SearchByUsername(username string) ([]*model.Hierarchy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filterCursor, err := hr.collection.Find(ctx, bson.M{"owner": username})
	if err != nil {
		return nil, err
	}

	var episodesFiltered []mongoHierarchy
	if err = filterCursor.All(ctx, &episodesFiltered); err != nil {
		return nil, err
	}

	return parseMongoMapToNativeHierarchies(episodesFiltered), nil
}
