package main

import (
	"github.com/abinashphulkonwar/redis/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Route("/api", routes.InsertHandler)

	app.Listen(":3000")
}
