package services

import (
	"context"
	"time"

	models "github.com/arihantdaga/kiotsundaython/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var jobsCollection = "scheudlejobsxzc"

func SaveJob(client *mongo.Client, job models.ScheduleJob) (interface{}, error) {
	collection := client.Database("kiotapp").Collection(jobsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, job)
	id := res.InsertedID
	return id, err
}

func EditJob(client *mongo.Client, job models.ScheduleJob) (interface{}, error) {
	collection := client.Database("kiotapp").Collection(jobsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, job, job) // TODO: Need to verify this.
	id := res.UpsertedID
	return id, err
}

func FindJob(client *mongo.Client, job models.ScheduleJob) (models.ScheduleJob, error) {
	collection := client.Database("kiotapp").Collection(jobsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var foundJob models.ScheduleJob
	defer cancel()
	err := collection.FindOne(ctx, bson.D{{"_id", job.ID}}).Decode(&foundJob)
	return foundJob, err
}
