package api

import (
	"github.com/abinashphulkonwar/redis/api/routes"
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	app := fiber.New()

	app.Route("/api/write", routes.InsertHandler)
	return app
}
