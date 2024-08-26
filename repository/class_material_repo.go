package repository

import (
    "context"
    "data-curation-test/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type ClassMaterialRepository interface {
    Create(ctx context.Context, classMaterial *models.ClassMaterial) error
    FindAll(ctx context.Context) ([]models.ClassMaterial, error)
    FindByID(ctx context.Context, id string) (*models.ClassMaterial, error)
    Update(ctx context.Context, id string, classMaterial *models.ClassMaterial) error
    Delete(ctx context.Context, id string) error
}

type classMaterialRepository struct {
    collection *mongo.Collection
}

func NewClassMaterialRepository(client *mongo.Client) ClassMaterialRepository {
    collection := client.Database("class").Collection("class_material")
    return &classMaterialRepository{collection: collection}
}

func (r *classMaterialRepository) Create(ctx context.Context, classMaterial *models.ClassMaterial) error {
    _, err := r.collection.InsertOne(ctx, classMaterial)
    return err
}

func (r *classMaterialRepository) FindAll(ctx context.Context) ([]models.ClassMaterial, error) {
    var classMaterials []models.ClassMaterial
    cursor, err := r.collection.Find(ctx, bson.D{{}})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var classMaterial models.ClassMaterial
        if err := cursor.Decode(&classMaterial); err != nil {
            return nil, err
        }
        classMaterials = append(classMaterials, classMaterial)
    }
    return classMaterials, nil
}

func (r *classMaterialRepository) FindByID(ctx context.Context, id string) (*models.ClassMaterial, error) {
    var classMaterial models.ClassMaterial
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&classMaterial)
    if err != nil {
        return nil, err
    }
    return &classMaterial, nil
}

func (r *classMaterialRepository) Update(ctx context.Context, id string, classMaterial *models.ClassMaterial) error {
    filter := bson.M{"_id": id}
    update := bson.M{"$set": classMaterial}
    _, err := r.collection.UpdateOne(ctx, filter, update)
    return err
}

func (r *classMaterialRepository) Delete(ctx context.Context, id string) error {
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}
