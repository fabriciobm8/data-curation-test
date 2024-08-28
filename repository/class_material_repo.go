package repository

import (
    "context"
    "data-curation-test/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "log"
    "errors"
)

type ClassMaterialRepository interface {
    Create(ctx context.Context, classMaterial *models.ClassMaterial) error
    FindAll(ctx context.Context) ([]models.ClassMaterial, error)
    FindByID(ctx context.Context, id string) (*models.ClassMaterial, error)
    Update(ctx context.Context, id string, classMaterial *models.ClassMaterial) error
    Delete(ctx context.Context, id string) error
    UpdateIsSuccessful(ctx context.Context, classMaterial *models.ClassMaterial, isSuccessful bool) error
    
}

type classMaterialRepository struct {
    collection *mongo.Collection
}

func NewClassMaterialRepository(client *mongo.Client) ClassMaterialRepository {
    collection := client.Database("class").Collection("classMaterial")
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

func (r *classMaterialRepository) UpdateIsSuccessful(ctx context.Context, classMaterial *models.ClassMaterial, isSuccessful bool) error {
    // Verifica se o documento existe com base no ID fornecido
    filter := bson.M{
        "_id": classMaterial.ID,  // Certifique-se de que o ID está sendo utilizado como chave
    }
    
    // Define o campo que será atualizado
    update := bson.M{
        "$set": bson.M{"isSuccessful": isSuccessful},
    }
    
    // Executa a atualização
    result, err := r.collection.UpdateOne(ctx, filter, update)
    if err != nil {
        log.Println("Error updating class material:", err.Error())
        return err
    }

    // Verifica se algum documento foi modificado
    if result.ModifiedCount == 0 {
        log.Println("No documents were updated")
        return errors.New("nenhum documento foi atualizado, verifique o ID fornecido")
    }

    log.Println("Updated class material successfully")
    return nil
}
