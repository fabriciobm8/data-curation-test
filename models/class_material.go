package models


type ClassMaterial struct {
    ID            string             `bson:"_id,omitempty"`
    UuidCourse    string             `bson:"uuidCourse"`
    UuidObjective string             `bson:"uuidObjective"`
    UuidMaterial  string             `bson:"uuidMaterial"`
    Transcript    string             `bson:"transcript,omitempty"`
    MaterialType  string             `bson:"materialType"`
    IsSuccessful  bool               `bson:"isSuccessful"`
}
