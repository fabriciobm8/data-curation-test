package repository

import (
    "context"
    "data-curation-test/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type TranscriptTimeRepository interface {
    Create(ctx context.Context, transcriptTime *models.TranscriptTime) error
    FindAll(ctx context.Context) ([]models.TranscriptTime, error)
    FindByID(ctx context.Context, id string) (*models.TranscriptTime, error)
    Update(ctx context.Context, id string, transcriptTime *models.TranscriptTime) error
    Delete(ctx context.Context, id string) error
}

type transcriptTimeRepository struct {
    collection *mongo.Collection
}

func NewTranscriptTimeRepository(client *mongo.Client) TranscriptTimeRepository {
    collection := client.Database("class").Collection("transcriptTime")
    return &transcriptTimeRepository{collection: collection}
}

func (r *transcriptTimeRepository) Create(ctx context.Context, transcriptTime *models.TranscriptTime) error {
    _, err := r.collection.InsertOne(ctx, transcriptTime)
    return err
}

func (r *transcriptTimeRepository) FindAll(ctx context.Context) ([]models.TranscriptTime, error) {
    var transcriptTimes []models.TranscriptTime
    cursor, err := r.collection.Find(ctx, bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var transcriptTime models.TranscriptTime
        if err := cursor.Decode(&transcriptTime); err != nil {
            return nil, err
        }
        transcriptTimes = append(transcriptTimes, transcriptTime)
    }
    return transcriptTimes, nil
}

func (r *transcriptTimeRepository) FindByID(ctx context.Context, id string) (*models.TranscriptTime, error) {
    var transcriptTime models.TranscriptTime
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&transcriptTime)
    if err != nil {
        return nil, err
    }
    return &transcriptTime, nil
}

func (r *transcriptTimeRepository) Update(ctx context.Context, id string, transcriptTime *models.TranscriptTime) error {
    filter := bson.M{"_id": id}
    update := bson.M{"$set": transcriptTime}
    _, err := r.collection.UpdateOne(ctx, filter, update)
    return err
}

func (r *transcriptTimeRepository) Delete(ctx context.Context, id string) error {
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}
