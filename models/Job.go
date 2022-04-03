package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScheduleJob struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	JobType   string             `bson:"jobtype"`
	JobMeta   string             `bson:"jobMeta"`
	JobTime   time.Time          `bson:"jobTime"`
	JobStatus string             `bson:"jobStatus"`
	JockedAt  time.Time          `bson:"lockedAt"`
}
