package services

import (
	"context"
	"time"

	models "github.com/arihantdaga/kiotsundaython/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var jobsCollection = "scheudlejobsxzc"

func SaveJob(client *mongo.Client, job models.ScheduleJob) (interface{}, error) {
	collection := client.Database("kiotapp").Collection(jobsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	job.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, job)
	if err != nil {
		return 0, err
	}
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

func FindJobsToBeExecuted(client *mongo.Client, ctx context.Context) ([]models.ScheduleJob, error) {
	collection := client.Database("kiotapp").Collection(jobsCollection)
	now := time.Now()
	var foundJobs []models.ScheduleJob
	cursor, err := collection.Find(ctx, bson.D{
		{"jobTime", bson.D{{"$lt", now}}},
	})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &foundJobs); err != nil {
		return nil, err
	}
	return foundJobs, nil
}

func LockJobs(client *mongo.Client, ctx context.Context, jobIds []primitive.ObjectID) error {
	collection := client.Database("kiotapp").Collection(jobsCollection)
	now := time.Now()
	_, err := collection.UpdateMany(ctx, bson.D{
		{"_id", bson.D{
			{"$in", jobIds},
		}},
	}, bson.D{
		{"$set", bson.D{
			{"lockedAt", now},
			{"jobStatus", "locked"},
		}},
	})
	if err != nil {
		return err
	}
	return nil
}
