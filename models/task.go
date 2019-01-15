package models

// Task is the structure for task
type Task struct {
	ID          string `json:"id" bson:"id"`
	ProjectKey  string `json:"projectKey" bson:"projectKey"`
	Summary     string `json:"summary" bson:"summary"`
	Description string `json:"description" bson:"description"`
}
