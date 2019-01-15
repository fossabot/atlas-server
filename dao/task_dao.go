package dao

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/atlas-io/atlas-server/client"

	"github.com/atlas-io/atlas-server/models"
)

// PutTask saves a new task to the mongo database
func PutTask(task *models.Task, onSuccess func(), onError func(error)) {
	mDb := client.GetMongoDefaultDatabase()
	collection := mDb.Collection("tasks")

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		log.Println(err.Error())
		onError(err)
	} else {
		onSuccess()
	}
}

// GetTask fetches a task from the mongo database
func GetTask(id string, onSuccess func(*models.Task), onError func(error)) {
	mDb := client.GetMongoDefaultDatabase()
	collection := mDb.Collection("tasks")

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	task := new(models.Task)
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&task)
	if err != nil {
		log.Println(err.Error())
		onError(err)
	} else {
		onSuccess(task)
	}
}
