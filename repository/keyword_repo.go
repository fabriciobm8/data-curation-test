package repository

import (
    "context"
    "data-curation-test/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type KeywordRepository interface {
    Create(ctx context.Context, keyword *models.Keyword) error
    FindAll(ctx context.Context) ([]models.Keyword, error)
    FindByID(ctx context.Context, id string) (*models.Keyword, error)
    Update(ctx context.Context, id string, keyword *models.Keyword) error
    Delete(ctx context.Context, id string) error
}

type keywordRepository struct {
    collection *mongo.Collection
}

func NewKeywordRepository(client *mongo.Client) KeywordRepository {
    collection := client.Database("class").Collection("keyword")
    return &keywordRepository{collection: collection}
}

func (r *keywordRepository) Create(ctx context.Context, keyword *models.Keyword) error {
    _, err := r.collection.InsertOne(ctx, keyword)
    return err
}

func (r *keywordRepository) FindAll(ctx context.Context) ([]models.Keyword, error) {
    var keywords []models.Keyword
    cursor, err := r.collection.Find(ctx, bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var keyword models.Keyword
        if err := cursor.Decode(&keyword); err != nil {
            return nil, err
        }
        keywords = append(keywords, keyword)
    }
    return keywords, nil
}

func (r *keywordRepository) FindByID(ctx context.Context, id string) (*models.Keyword, error) {
    var keyword models.Keyword
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&keyword)
    if err != nil {
        return nil, err
    }
    return &keyword, nil
}

func (r *keywordRepository) Update(ctx context.Context, id string, keyword *models.Keyword) error {
    filter := bson.M{"_id": id}
    update := bson.M{"$set": keyword}
    _, err := r.collection.UpdateOne(ctx, filter, update)
    return err
}

func (r *keywordRepository) Delete(ctx context.Context, id string) error {
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}
