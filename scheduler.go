package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/arihantdaga/kiotsundaython/database"
	"github.com/arihantdaga/kiotsundaython/runner"
	"github.com/joho/godotenv"
)

func main() {
	println("Starting scheduler...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBStr := os.Getenv("DB_CONNECTION")
	dbClient, err := database.DB(DBStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = dbClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second * 3)
	if err := dbClient.Ping(context.TODO(), nil); err != nil {
		log.Println("Error pinging database. Exiting...")
		panic(err)
	} else {
		log.Println("Successfully pinged database")
	}

	// Runner
	runnerDie := make(chan string, 1)
	runner := runner.JobRunnerImpl{}
	runner.New(dbClient)
	go runner.Run(runnerDie)

	// time.Sleep(time.Second * 10)

	// API Setup
	// TODO:I think this is not the right way - I think New should return something like a pointer to APIServerImpl.
	api := APiServerImpl{}
	api.New(dbClient)
	api.HandleRoutes()
	api.Start()

	print(<-runnerDie)
}
