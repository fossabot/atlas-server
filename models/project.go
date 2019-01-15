package models

// Project is the structure for projects
type Project struct {
	Key  string `json:"key" bson:"key"`
	Name string `json:"name" bson:"name"`
}
