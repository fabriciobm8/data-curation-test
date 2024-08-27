package models


type Keyword struct {
    ID           string             `bson:"_id,omitempty"`
    CourseId     string             `bson:"uuidCourse"`
    Keyword	  	 string             `bson:"uuidObjective"`
    UsageCount   int                `bson:"uuidMaterial"`
}
