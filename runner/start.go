package runner

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/arihantdaga/kiotsundaython/database"
	"github.com/joho/godotenv"
)

func rundummy() {
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

	// runner := JobRunnerImpl{}
	// runner.New(dbClient)
	// runner.Run()
}
