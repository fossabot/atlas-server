package dao

import (
	"context"
	"log"
	"time"

	"github.com/atlas-io/atlas-server/client"
	"github.com/atlas-io/atlas-server/models"
	"github.com/mongodb/mongo-go-driver/bson"
)

// PutProject saves a `project` to the mongo database.
/*
 * Use the `onSuccess` to handle succesful addition
 * Use the `onError` to capture and handle errors
 */
func PutProject(project *models.Project, onSuccess func(), onError func(error)) {
	mDb := client.GetMongoDefaultDatabase()
	collection := mDb.Collection("projects")

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	_, err := collection.InsertOne(ctx, project)
	if err != nil {
		log.Println(err.Error())
		onError(err)
	} else {
		onSuccess()
	}
}

// GetProject fetches a `project` from the mongo database.
/*
 * Use the `onSuccess` to handle succesful database fetch
 * Use the `onError` to capture and handle errors
 */
func GetProject(key string, onSuccess func(*models.Project), onError func(error)) {
	mDb := client.GetMongoDefaultDatabase()
	collection := mDb.Collection("projects")

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	project := new(models.Project)
	err := collection.FindOne(ctx, bson.M{"key": key}).Decode(&project)
	if err != nil {
		onError(err)
	} else {
		onSuccess(project)
	}
}
