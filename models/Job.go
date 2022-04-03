package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScheduleJob struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	jobtype   string             `bson:"jobtype"`
	jobMeta   string             `bson:"jobMeta"`
	jobTime   time.Time          `bson:"jobTime"`
	jobStatus string             `bson:"jobStatus"`
	lockedAt  time.Time          `bson:"lockedAt"`
}
