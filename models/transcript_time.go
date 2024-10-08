package models


type TranscriptTime struct {
    ID              string             `bson:"_id,omitempty"`
    StartTime       float64            `bson:"startTime"`
    EndTime		    float64            `bson:"endTime"`
    Transcript      string             `bson:"transcript,omitempty"`
    ClassMaterialId string             `bson:"classMaterialId"`
}
