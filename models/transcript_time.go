package models


type TranscriptTime struct {
    ID            string             `bson:"_id,omitempty"`
    StartTime     float64            `bson:"uuidCourse"`
    EndTime		  float64             `bson:"uuidObjective"`
    UuidMaterial  string             `bson:"uuidMaterial"`
    Transcript    string             `bson:"transcript,omitempty"`
}
