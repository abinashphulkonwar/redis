package api

import (
	"github.com/abinashphulkonwar/redis/api/routes"
	"github.com/abinashphulkonwar/redis/internalstorage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func App(queue *internalstorage.Queue) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Route("/api/write", routes.InsertHandler(queue))
	app.Route("/api/query", routes.QueryHandler(queue))

	return app
}
