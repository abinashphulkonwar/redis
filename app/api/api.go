package api

import (
	"github.com/abinashphulkonwar/redis/api/routes"
	"github.com/abinashphulkonwar/redis/storage"
	"github.com/gofiber/fiber/v2"
)

func App(queue *storage.Queue) *fiber.App {
	app := fiber.New()

	app.Route("/api/write", routes.InsertHandler(queue))
	return app
}
