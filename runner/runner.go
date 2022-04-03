package runner

import (
	"context"
	"fmt"

	"github.com/arihantdaga/kiotsundaython/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type JobRunner interface {
	Run()
}
type JobRunnerImpl struct {
	client *mongo.Client
}

func (j *JobRunnerImpl) New(client *mongo.Client) {
	j.client = client
}
func (j *JobRunnerImpl) Run() {
	ctx := context.TODO()
	fmt.Println("Job runner running...")
	jobs, err := services.FindJobsToBeExecuted(j.client, ctx)
	if err != nil {
		println("Could not find Jobs. WIll try next time")
		return
	}
	if len(jobs) > 0 {
		fmt.Printf("Found Jobs to execute: %d ", len(jobs))
	}
	jobIds := make([]primitive.ObjectID, len(jobs))
	for i := 0; i < len(jobs); i++ {
		jobIds[i] = jobs[i].ID
	}
	if err := services.LockJobs(j.client, ctx, jobIds); err != nil {
		fmt.Println("Could not lock jobs")
	}
	for _, j := range jobs {
		// Call webhook.
		println(j.JobMeta)
	}
}
