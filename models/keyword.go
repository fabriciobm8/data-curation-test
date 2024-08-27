package models


type Keyword struct {
    ID                  string             `bson:"_id,omitempty"`
    CourseId            string             `bson:"courseId"`
    Keyword	  	        string             `bson:"keyword"`
    UsageCount          int                `bson:"usageCount"`
    ClassMaterialId     string             `bson:"classMaterialId"`
    TranscriptTimeId    string             `bson:"transcriptTimeId"`
}
