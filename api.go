package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type APiServer interface {
	New()
}
type APiServerImpl struct {
	client *mongo.Client
	app    *fiber.App
}

func (a *APiServerImpl) New(client *mongo.Client) {
	a.client = client
	a.app = fiber.New()
}

func (a *APiServerImpl) Start() {
	a.app.Listen(":8080")
}

func (a *APiServerImpl) HandleRoutes() {
	app := a.app
	app.Get("/api/v1/jobs", func(c *fiber.Ctx) error {
		return c.SendString("all jobs")
	})
}
